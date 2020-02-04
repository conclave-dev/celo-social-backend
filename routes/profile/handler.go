package profile

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// AddRoutes includes routes related to the `celo` endpoints on the RPC
func AddRoutes(router *mux.Router) {
	router.HandleFunc("/profile/update", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("/profile/update placeholder")
	})
}
