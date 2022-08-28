package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/commit-smart-core-banking-system/config"
	"github.com/commit-smart-core-banking-system/model"
	"github.com/commit-smart-core-banking-system/response/http_response"
	"github.com/commit-smart-core-banking-system/response/message"
	"github.com/golang-jwt/jwt/v4"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {

		if request.Header["Authorization"] == nil {
			http_response.ErrorResponse(response, http.StatusUnauthorized, message.ResponseMessage.Unauthorized)
			return
		}

		token, err := jwt.Parse(strings.Split(request.Header["Authorization"][0], "Bearer ")[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("")
			}
			return config.SECRET_KEY, nil
		})

		if err != nil {
			http_response.ErrorResponse(response, http.StatusBadRequest, message.ResponseMessage.TokenExpried)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			accountId := claims[model.Claim.AccountId]
			request.Header.Set(model.Claim.AccountId, fmt.Sprintf("%f", accountId))
			handler.ServeHTTP(response, request)
			return
		}
		http_response.ErrorResponse(response, http.StatusUnauthorized, message.ResponseMessage.Unauthorized)
		return
	}
}
