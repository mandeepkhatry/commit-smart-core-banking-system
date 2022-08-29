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

func Transfer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var transferRequest model.TransferRequest
	json.NewDecoder(request.Body).Decode(&transferRequest)

	//Validate Request Body
	validateResp := validator.ValidateRequest(transferRequest)
	if !validateResp.Validate {
		http_response.ErrorResponse(response, http.StatusBadRequest, validateResp.Errors)
		return
	}

	//Get AccountId From Token and Validate Sernder's Account
	transferRequest.AccountId = GetAccountIdFromRequest(request)
	senderAccountResponse := ValidateAccount(transferRequest.AccountId)
	if !senderAccountResponse.Validate {
		http_response.ErrorResponse(response, http.StatusNotFound, message.ResponseMessage.ItemNotFound)
		return
	}

	if transferRequest.AccountId == transferRequest.CorrespondingAccountId {
		http_response.ErrorResponse(response, http.StatusBadRequest, message.ResponseMessage.InvalidTransaction)
		return
	}

	//Validate receiver's account
	receiverAccountResponse := ValidateAccount(transferRequest.CorrespondingAccountId)
	if !receiverAccountResponse.Validate {
		http_response.ErrorResponse(response, http.StatusNotFound, message.ResponseMessage.ItemNotFound)
		return
	}

	//Update Account Balance
	serializedAmount := utils.SerializeAmount(transferRequest.Amount)
	senderAccountResponse.Account.Balance -= serializedAmount
	if senderAccountResponse.Account.Balance < 0 {
		http_response.ErrorResponse(response, http.StatusBadRequest, message.ResponseMessage.InsufficientFund)
		return
	}

	senderUpdatedAccountResponse, err := store.Store.UpdateAccount(senderAccountResponse.Account)
	if err != nil {
		logger.Info(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	var senderTxnRequest = model.TxnRequest{}
	senderTxnRequest.AccountId = transferRequest.AccountId
	senderTxnRequest.TransactionType = model.TransactionType.Transfer
	senderTxnRequest.AmountType = model.AmountType.DR
	senderTxnRequest.Amount = serializedAmount
	senderTxnRequest.Description = transferRequest.Description
	senderTxnRequest.CorrespondingAccountId = transferRequest.CorrespondingAccountId
	senderTxnRequest.CurrentBalance = senderUpdatedAccountResponse.Balance

	senderTxnResponse, err := store.Store.CreateTxn(senderTxnRequest)
	if err != nil {
		logger.Info(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	receiverAccountResponse.Account.Balance += serializedAmount

	receiverUpdatedAccountResponse, err := store.Store.UpdateAccount(receiverAccountResponse.Account)
	if err != nil {
		logger.Info(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	var receiverTxnRequest = model.TxnRequest{}
	receiverTxnRequest.AccountId = transferRequest.CorrespondingAccountId
	receiverTxnRequest.TransactionType = model.TransactionType.Transfer
	receiverTxnRequest.AmountType = model.AmountType.CR
	receiverTxnRequest.Amount = serializedAmount
	receiverTxnRequest.Description = transferRequest.Description
	receiverTxnRequest.CorrespondingAccountId = transferRequest.AccountId
	receiverTxnRequest.CurrentBalance = receiverUpdatedAccountResponse.Balance

	_, err = store.Store.CreateTxn(receiverTxnRequest)
	if err != nil {
		logger.Info(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	http_response.Response(response, http.StatusCreated, DeserializedTxns(model.ListTxnResponse{Txn: []model.TxnResponse{senderTxnResponse}}).Txn[0])
}
