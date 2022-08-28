package api

import (
	"net/http"

	"github.com/commit-smart-core-banking-system/logger"
	"github.com/commit-smart-core-banking-system/model"
	"github.com/commit-smart-core-banking-system/response/http_response"
	"github.com/commit-smart-core-banking-system/response/message"
	"github.com/commit-smart-core-banking-system/store"
	"github.com/commit-smart-core-banking-system/utils"
	"github.com/commit-smart-core-banking-system/validator"
	"github.com/gorilla/schema"
)

func ListTransaction(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	if err := request.ParseForm(); err != nil {
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	var txnFilter model.TransactionFilter
	if err := schema.NewDecoder().Decode(&txnFilter, request.Form); err != nil {
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}
	validateResp := validator.ValidateRequest(txnFilter)
	if !validateResp.Validate {
		http_response.ErrorResponse(response, http.StatusBadRequest, validateResp.Errors)
		return
	}

	accountId := GetAccountIdFromRequest(request)
	accountResponse := ValidateAccount(accountId)
	if !accountResponse.Validate {
		http_response.ErrorResponse(response, http.StatusNotFound, message.ResponseMessage.ItemNotFound)

	}

	txnFilter.Offset = utils.GetOffset(txnFilter.Page, txnFilter.Limit)
	txnFilter.AccountId = accountResponse.Account.Id

	accountList, err := store.Store.ListTxn(txnFilter)
	if err != nil {
		logger.Error(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	customerTransactions := DeserializedTxns(accountList)
	customerTransactions.Balance = utils.DeserializeAmount(accountResponse.Account.Balance)

	http_response.Response(response, http.StatusCreated, customerTransactions)
}

func DeserializedTxns(list model.ListTxnResponse) model.ListTransactionResponse {
	var txnList model.ListTransactionResponse
	for _, txn := range list.Txn {
		var deserializedTxn = model.TransactionResponse{
			MetaInfo:       txn.MetaInfo,
			TxnInfo:        txn.TxnInfo,
			Amount:         utils.DeserializeAmount(txn.Amount),
			CurrentBalance: utils.DeserializeAmount(txn.CurrentBalance),
		}
		txnList.Txn = append(txnList.Txn, deserializedTxn)
	}

	return txnList

}
