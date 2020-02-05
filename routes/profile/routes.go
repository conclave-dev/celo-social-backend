package profile

import (
	"github.com/gorilla/mux"
)

// AddRoutes includes routes related to the `celo` endpoints on the RPC
func AddRoutes(router *mux.Router) {
	router.HandleFunc("/profile/update", updateProfile)
}
