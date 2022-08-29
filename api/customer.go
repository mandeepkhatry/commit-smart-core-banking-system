package api

import (
	"encoding/json"
	"net/http"

	"github.com/commit-smart-core-banking-system/logger"
	"github.com/commit-smart-core-banking-system/model"
	"github.com/commit-smart-core-banking-system/response/http_response"
	"github.com/commit-smart-core-banking-system/response/message"
	"github.com/commit-smart-core-banking-system/store"
	"github.com/commit-smart-core-banking-system/utils/password"
	"github.com/commit-smart-core-banking-system/utils/token"
	"github.com/commit-smart-core-banking-system/validator"
)

func CustomerRegister(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var customerRequest model.CustomerRequest
	json.NewDecoder(request.Body).Decode(&customerRequest)

	//Validate Request Body
	validateResp := validator.ValidateRequest(customerRequest)
	if !validateResp.Validate {
		http_response.ErrorResponse(response, http.StatusBadRequest, validateResp.Errors)
		return
	}

	_, err := store.Store.FetchCustomer(customerRequest.EmailAddress)
	if err == nil {
		http_response.ErrorResponse(response, http.StatusBadRequest, message.ResponseMessage.CustomerAlreadyExists)
		return
	}

	passwordHashResp := password.HashPassword(customerRequest.Password)
	if passwordHashResp.Error != nil {
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}
	customerRequest.PasswordHash = passwordHashResp.HashedPassword

	customerResponse, err := store.Store.CreateCustomer(customerRequest.Customer)
	if err != nil {
		logger.Debug(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	_, err = store.Store.CreateAccount(model.Account{
		Type:       model.AccountType.Customer,
		CustomerId: customerResponse.Id,
		Balance:    0,
	})

	if err != nil {
		logger.Debug(err.Error())
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	http_response.Response(response, http.StatusCreated, customerResponse)
}

func CustomerLogin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var loginData model.Login
	json.NewDecoder(request.Body).Decode(&loginData)

	validateResp := validator.ValidateRequest(loginData)
	if !validateResp.Validate {
		http_response.ErrorResponse(response, http.StatusBadRequest, validateResp.Errors)
		return
	}

	customerData, err := store.Store.FetchCustomer(loginData.Email)
	if err != nil {
		http_response.ErrorResponse(response, http.StatusNotFound, message.ResponseMessage.ItemNotFound)
		return
	}

	if err = password.ValidatePassword(password.ValidatePasswordParams{
		Password:       loginData.Password,
		HashedPassword: customerData.PasswordHash,
	}); err != nil {
		http_response.ErrorResponse(response, http.StatusUnauthorized, message.ResponseMessage.Unauthorized)
		return
	}

	customerAccount, err := store.Store.FetchAccountByCustomerId(customerData.Id)
	if err != nil {
		http_response.ErrorResponse(response, http.StatusNotFound, message.ResponseMessage.ItemNotFound)
		return
	}

	tokenResponse := token.GenerateJWT(token.TokenParams{
		Email:     customerData.EmailAddress,
		AccountId: customerAccount.Id,
	})

	if tokenResponse.Error != nil {
		http_response.ErrorResponse(response, http.StatusInternalServerError, message.ResponseMessage.SomethingWentWrong)
		return
	}

	http_response.Response(response, http.StatusOK, model.TokenResponse{Token: tokenResponse.Token})
}
