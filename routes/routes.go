package routes

import (
	"github.com/gorilla/mux"
	"github.com/stella-zone/celo-social-backend/routes/profile"
)

// SetUpRoutes initiates the setup process for all routes
func SetUpRoutes(router *mux.Router) {
	profile.AddRoutes(router)
}
