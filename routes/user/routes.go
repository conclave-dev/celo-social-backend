package user

import (
	"github.com/gorilla/mux"
)

// AddRoutes includes routes related to the `celo` endpoints on the RPC
func AddRoutes(router *mux.Router) {
	router.HandleFunc(`/user/{user}`, getUser)
	router.HandleFunc(`/user/{user}/update`, updateUser)
}
