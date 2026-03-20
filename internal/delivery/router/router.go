package router

import (
	"github.com/gorilla/mux"
	"visualizationBdDebet/internal/delivery/handler"
)

// NewRouter создаёт новый роутер и регистрирует все маршруты приложения.
func NewRouter(
	debetHandler *handler.DebetHandler,
	contractHandler *handler.ContractHandler,
	blockfactorHandler *handler.BlockFactorHandler,
) *mux.Router {
	r := mux.NewRouter()

	// Группа маршрутов для debet
	debet := r.PathPrefix("/debet").Subrouter()
	debet.HandleFunc("", debetHandler.GetAll).Methods("GET")
	debet.HandleFunc("/{orgName}", debetHandler.GetByOrgName).Methods("GET")

	contract := r.PathPrefix("/contract").Subrouter()
	contract.HandleFunc("", contractHandler.GetAll).Methods("GET")
	contract.HandleFunc("/{id}", contractHandler.GetById).Methods("GET")

	blockfactor := r.PathPrefix("/blockfactor").Subrouter()
	blockfactor.HandleFunc("", blockfactorHandler.GetAll).Methods("GET")
	blockfactor.HandleFunc("/{orgName}", blockfactorHandler.GetById).Methods("GET")

	// Middleware можно добавить глобально или на подроутер
	// r.Use(middleware.Logger, middleware.Recoverer)

	return r
}
