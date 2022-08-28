package model

var AccountType = struct {
	Bank     string
	Customer string
}{
	"bank",
	"customer",
}

type AccountInfo struct {
	Type       string `json:"type" validate:"oneof=bank customer"`
	CustomerId int64  `json:"customer_id" validate:"required"`
	Balance    int64  `json:"balance"`
}

type AccountRequest struct {
	Account
}

type AccountResponse struct {
	MetaInfo
	AccountInfo
}

type AccountFilter struct {
	Type string `schema:"type" json:"type"  validate:"required"`
	Filter
}

type AccountCustomer struct {
	MetaInfo
	Type string `json:"type"`
	Name string `json:"customer_name"`
}

type ListAccountResponse struct {
	Accounts []AccountCustomer `json:"data"`
}

type Account struct {
	Balance    int64  `json:"balance"`
	Type       string `json:"type" validate:"oneof=bank customer"`
	CustomerId int64  `json:"customer_id" validate:"required"`
}
