package model

type Txn struct {
	TxnInfo
	Amount int64 `json:"amount" validate:"required,numeric,gt=0"`
	CorrespondingAccount
	CurrentBalance int64 `json:"current_balance"`
}

type CorrespondingAccount struct {
	CorrespondingAccountId int64 `json:"corresponding_account_id"`
}

//Deposit Withdraw
type TransactionRequest struct {
	TxnInfo
	CurrentBalance int64   `json:"current_balance"`
	Amount         float64 `json:"amount" validate:"required,numeric,gt=0"`
}

//Transfer
type TransferRequest struct {
	TxnInfo
	Amount float64 `json:"amount" validate:"required,numeric,gt=0"`
	CorrespondingAccount
	CurrentBalance int64 `json:"current_balance"`
}

type TxnRequest struct {
	Txn
}

type TxnResponse struct {
	MetaInfo
	Txn
}

type TxnInfo struct {
	AccountId       int64  `json:"account_id"`
	TransactionType string `json:"transaction_type"`
	AmountType      string `json:"amount_type"`
	Description     string `json:"description" validate:""`
}

var AmountType = struct {
	DR string
	CR string
}{
	"dr",
	"cr",
}

var TransactionType = struct {
	Deposit  string
	Withdraw string
	Transfer string
}{
	"deposit",
	"withdraw",
	"transfer",
}

type ListTxnResponse struct {
	Txn []TxnResponse `json:"data"`
}

//Response to API CALL
type TransactionResponse struct {
	MetaInfo
	TxnInfo
	CurrentBalance float64 `json:"current_balance"`
	Amount         float64 `json:"amount" validate:"required,numeric,gt=0"`
}

type ListTransactionResponse struct {
	Txn     []TransactionResponse `json:"data"`
	Balance float64               `json:"balance"`
}
