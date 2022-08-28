package api

import (
	"encoding/json"
	"net/http"

	"github.com/commit-smart-core-banking-system/logger"
	"github.com/commit-smart-core-banking-system/model"
	"github.com/commit-smart-core-banking-system/response/http_response"
	"github.com/commit-smart-core-banking-system/response/message"
	"github.com/commit-smart-core-banking-system/store"
	"github.com/commit-smart-core-banking-system/utils"
	"github.com/commit-smart-core-banking-system/validator"
)

func Withdraw(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var withdrawRequest model.TransactionRequest
	json.NewDecoder(request.Body).Decode(&withdrawRequest)

	validateResp := validator.ValidateRequest(withdrawRequest)
	if !validateResp.Validate {
		http_response.ErrorResponse(response, http.StatusBadRequest, validateResp.Errors)
		return
	}

	withdrawRequest.AccountId = GetAccountIdFromRequest(request)

	accountResponse := ValidateAccount(withdrawRequest.AccountId)

	if !accountResponse.Validate {
		http_response.ErrorResponse(response, http.StatusNotFound, message.ResponseMessage.ItemNotFound)
		return
	}

	//Update Account Balance
	serializedAmount := utils.SerializeAmount(withdrawRequest.Amount)
	accountResponse.Account.Balance -= serializedAmount

	if accountResponse.Account.Balance < 0 {
		http_response.ErrorResponse(response, http.StatusBadRequest, message.ResponseMessage.InsufficientFund)
		return
	}

	updatedAccountResponse, err := store.Store.UpdateAccount(accountResponse.Account)
	if err != nil {
		logger.Info(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	withdrawRequest.AmountType = model.AmountType.DR
	withdrawRequest.TransactionType = model.TransactionType.Withdraw
	withdrawRequest.CurrentBalance = updatedAccountResponse.Balance

	//Database compatible object
	var withdrawTxnRequest = model.TxnRequest{}
	withdrawTxnRequest.TxnInfo = withdrawRequest.TxnInfo
	withdrawTxnRequest.Amount = serializedAmount
	withdrawTxnRequest.CurrentBalance = withdrawRequest.CurrentBalance

	txnResponse, err := store.Store.CreateTxn(withdrawTxnRequest)
	if err != nil {
		logger.Info(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	http_response.Response(response, http.StatusCreated, DeserializedTxns(model.ListTxnResponse{Txn: []model.TxnResponse{txnResponse}}).Txn[0])
}
