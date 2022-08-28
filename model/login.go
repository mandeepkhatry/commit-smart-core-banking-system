package model

type Login struct {
	Email    string `json:"email_address" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type TokenResponse struct {
	Token string `jsom:"token"`
}
