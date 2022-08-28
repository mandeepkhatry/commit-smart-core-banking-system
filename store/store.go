package store

import (
	"github.com/commit-smart-core-banking-system/model"
)

type DataStore interface {
	Connection
	CreateCustomer(model.Customer) (model.CustomerResponse, error)
	FetchCustomer(email string) (model.CustomerData, error)

	CreateAccount(model.Account) (model.AccountResponse, error)
	UpdateAccount(model.AccountResponse) (model.AccountResponse, error)

	ListAccount(model.AccountFilter) (model.ListAccountResponse, error)
	FetchAccount(id int64) (model.AccountResponse, error)
	FetchAccountByCustomerId(id int64) (model.AccountResponse, error)

	CreateTxn(model.TxnRequest) (model.TxnResponse, error)
	ListTxn(model.TransactionFilter) (model.ListTxnResponse, error)
}

type Connection interface {
	CloseClient() error
}
