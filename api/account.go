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

func ListAccount(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	if err := request.ParseForm(); err != nil {
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	var filter model.AccountFilter
	if err := schema.NewDecoder().Decode(&filter, request.Form); err != nil {
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	//Validate Request Body
	validateResp := validator.ValidateRequest(filter)
	if !validateResp.Validate {
		http_response.ErrorResponse(response, http.StatusBadRequest, validateResp.Errors)
		return
	}

	filter.Offset = utils.GetOffset(filter.Page, filter.Limit)

	accountList, err := store.Store.ListAccount(filter)
	if err != nil {
		logger.Error(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	http_response.Response(response, http.StatusOK, accountList)
}

type ValidateAccountResp struct {
	Account  model.AccountResponse
	Validate bool
}

func ValidateAccount(id int64) ValidateAccountResp {
	accountResp, err := store.Store.FetchAccount(id)
	if err != nil {
		return ValidateAccountResp{Validate: false}
	}
	return ValidateAccountResp{Account: accountResp, Validate: true}
}
