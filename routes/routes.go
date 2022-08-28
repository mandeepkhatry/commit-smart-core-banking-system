package routes

import (
	"net/http"

	"github.com/commit-smart-core-banking-system/api"
	"github.com/commit-smart-core-banking-system/middleware"
	"github.com/gorilla/mux"
)

var HTTPMethod = struct {
	POST string
	GET  string
}{
	"POST",
	"GET",
}

//Register All Routes
func RegisterRoutes() http.Handler {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()

	r.HandleFunc("/register", api.CustomerRegister).Methods(HTTPMethod.POST)
	r.HandleFunc("/login", api.CustomerLogin).Methods(HTTPMethod.POST)

	r.HandleFunc("/deposit", middleware.IsAuthorized(api.Deposit)).Methods(HTTPMethod.POST)
	r.HandleFunc("/withdraw", middleware.IsAuthorized(api.Withdraw)).Methods(HTTPMethod.POST)
	r.HandleFunc("/transfer", middleware.IsAuthorized(api.Transfer)).Methods(HTTPMethod.POST)
	r.HandleFunc("/transactions", middleware.IsAuthorized(api.ListTransaction)).Methods(HTTPMethod.GET)

	r.HandleFunc("/accounts", middleware.IsAuthorized(api.ListAccount)).Methods(HTTPMethod.GET)

	return r
}
