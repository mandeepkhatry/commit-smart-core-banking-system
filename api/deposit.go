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

func Deposit(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var depositRequest model.TransactionRequest
	json.NewDecoder(request.Body).Decode(&depositRequest)

	//Validate Request Body
	validateResp := validator.ValidateRequest(depositRequest)
	if !validateResp.Validate {
		http_response.ErrorResponse(response, http.StatusBadRequest, validateResp.Errors)
		return
	}

	//Get AccountId From Token and Validate Account
	depositRequest.AccountId = GetAccountIdFromRequest(request)
	accountResponse := ValidateAccount(depositRequest.AccountId)

	if !accountResponse.Validate {
		http_response.ErrorResponse(response, http.StatusNotFound, message.ResponseMessage.ItemNotFound)
		return
	}

	//Update Account Balance
	serializedAmount := utils.SerializeAmount(depositRequest.Amount)
	accountResponse.Account.Balance += serializedAmount

	updatedAccountResponse, err := store.Store.UpdateAccount(accountResponse.Account)
	if err != nil {
		logger.Info(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	depositRequest.AmountType = model.AmountType.CR
	depositRequest.TransactionType = model.TransactionType.Deposit
	depositRequest.CurrentBalance = updatedAccountResponse.Balance

	//Database compatible object
	var depositTxnRequest = model.TxnRequest{}
	depositTxnRequest.TxnInfo = depositRequest.TxnInfo
	depositTxnRequest.Amount = serializedAmount
	depositTxnRequest.CurrentBalance = depositRequest.CurrentBalance

	txnResponse, err := store.Store.CreateTxn(depositTxnRequest)
	if err != nil {
		logger.Info(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	http_response.Response(response, http.StatusCreated, DeserializedTxns(model.ListTxnResponse{Txn: []model.TxnResponse{txnResponse}}).Txn[0])
}

func GetAccountIdFromRequest(request *http.Request) int64 {
	return utils.StringToInt(request.Header.Get(model.Claim.AccountId))
}
