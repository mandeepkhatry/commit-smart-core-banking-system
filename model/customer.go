package model

type CustomerInfo struct {
	Name         string `json:"name" validate:"required"`
	Address      string `json:"address" validate:"required"`
	PhoneNumber  int    `json:"phone_number" validate:"required,numeric,min=10"`
	EmailAddress string `json:"email_address" validate:"required,email"`
}

type Customer struct {
	CustomerInfo
	PasswordHash string `json:"password_hash"`
}

type CustomerRequest struct {
	Customer
	Password string `json:"password" validate:"required,min=8"`
}

type CustomerResponse struct {
	MetaInfo
	CustomerInfo
}

type CustomerData struct {
	MetaInfo
	Customer
}
