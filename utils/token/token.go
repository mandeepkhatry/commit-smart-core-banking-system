package token

import (
	"time"

	"github.com/commit-smart-core-banking-system/config"
	"github.com/commit-smart-core-banking-system/model"
	"github.com/golang-jwt/jwt"
)

type TokenParams struct {
	Email     string
	AccountId int64
}

type TokenResponse struct {
	Token string
	Error error
}

var tokenConfig = struct {
	UnitInMinutes int
}{
	24,
}

func GenerateJWT(params TokenParams) TokenResponse {
	var mySigningKey = config.SECRET_KEY
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[model.Claim.Authorized] = true
	claims[model.Claim.Email] = params.Email
	claims[model.Claim.AccountId] = params.AccountId
	claims[model.Claim.Exp] = time.Now().Add(time.Hour * time.Duration(tokenConfig.UnitInMinutes)).Unix()
	tokenString, err := token.SignedString(mySigningKey)
	return TokenResponse{Token: tokenString, Error: err}

}
