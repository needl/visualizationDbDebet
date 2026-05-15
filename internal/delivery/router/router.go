package router

import (
	"visualizationBdDebet/internal/blockfactor"
	"visualizationBdDebet/internal/contract"
	"visualizationBdDebet/internal/contractor"
	"visualizationBdDebet/internal/customer"
	"visualizationBdDebet/internal/debet"
	"visualizationBdDebet/internal/object"
	"visualizationBdDebet/internal/response"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router and registers all application routes.
func NewRouter(
	debetHandler *debet.Handler,
	contractHandler *contract.Handler,
	blockfactorHandler *blockfactor.Handler,
	responseHandler *response.Handler,
	customerHandler *customer.Handler,
	contractorHandler *contractor.Handler,
	objectHandler *object.Handler,
) *mux.Router {
	r := mux.NewRouter()

	debet.RegisterRoutes(r, debetHandler)
	contract.RegisterRoutes(r, contractHandler)
	blockfactor.RegisterRoutes(r, blockfactorHandler)
	response.RegisterRoutes(r, responseHandler)
	customer.RegisterRoutes(r, customerHandler)
	contractor.RegisterRoutes(r, contractorHandler)
	object.RegisterRoutes(r, objectHandler)

	return r
}
