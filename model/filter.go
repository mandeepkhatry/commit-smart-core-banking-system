package model

/*
	Set of all filters
*/

type Filter struct {
	Limit  int64 `schema:"limit" json:"limit"  validate:"required"`
	Page   int64 `schema:"page" json:"page"  validate:"required"`
	Offset int64
}

type TransactionFilter struct {
	Type      string `schema:"type" json:"type" validate:"oneof=deposit withdraw all"`
	Date      string `schema:"date" json:"date"`
	Order     string `schema:"order" json:"order" validate:"oneof=desc asc"`
	AccountId int64
	Filter
}

var FilterParam = struct {
	All string
}{
	"all",
}

var Claim = struct {
	AccountId  string
	Email      string
	Authorized string
	Exp        string
}{
	"account_id",
	"email",
	"authorized",
	"exp",
}
