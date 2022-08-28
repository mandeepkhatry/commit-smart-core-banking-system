package message

var ResponseMessage = struct {
	SomethingWentWrong    string
	ItemNotFound          string
	InsufficientFund      string
	Unauthorized          string
	TokenExpried          string
	CustomerAlreadyExists string
	InvalidTransaction    string
}{
	"Something Went Wrong",
	"Not Found",
	"Insufficient Fund",
	"Unauthorized",
	"Token Expired",
	"Customer Already Exists",
	"Invalid Transaction",
}
