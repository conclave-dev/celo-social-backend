package routes

import (
	"github.com/gorilla/mux"
	"github.com/stella-zone/celo-social-backend/routes/profile"
	"github.com/stella-zone/celo-social-backend/routes/service"
)

// SetUpRoutes initiates the setup process for all routes
func SetUpRoutes(router *mux.Router) {
	profile.AddRoutes(router)
	service.AddRoutes(router)
}
