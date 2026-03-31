package router

import (
	"github.com/gorilla/mux"
	"visualizationBdDebet/internal/delivery/customer"
	"visualizationBdDebet/internal/delivery/handler"
	"visualizationBdDebet/internal/delivery/response"
)

// NewRouter создаёт новый роутер и регистрирует все маршруты приложения.
func NewRouter(
	debetHandler *handler.DebetHandler,
	contractHandler *handler.ContractHandler,
	blockfactorHandler *handler.BlockFactorHandler,
	responseHandler *response.Handler,
	customerHandler *customer.Handler,
) *mux.Router {
	r := mux.NewRouter()

	// Группа маршрутов для debet
	debet := r.PathPrefix("/debet").Subrouter()
	debet.HandleFunc("", debetHandler.GetAll).Methods("GET")
	debet.HandleFunc("/withMIP", debetHandler.GetAllWithMIP).Methods("GET")

	contract := r.PathPrefix("/contract").Subrouter()
	contract.HandleFunc("", contractHandler.GetAll).Methods("GET")
	contract.HandleFunc("/{id}", contractHandler.GetById).Methods("GET")

	blockfactor := r.PathPrefix("/blockFactor").Subrouter()
	blockfactor.HandleFunc("", blockfactorHandler.GetAll).Methods("GET")
	blockfactor.HandleFunc("/{orgName}", blockfactorHandler.GetById).Methods("GET")

	resp := r.PathPrefix("/response").Subrouter()
	resp.HandleFunc("", responseHandler.GetResponse).Methods("GET")
	resp.HandleFunc("/withMIP", responseHandler.GetResponseWithMIP).Methods("GET")

	customer := r.PathPrefix("/customer").Subrouter()
	customer.HandleFunc("", customerHandler.GetCustomers).Methods("GET")
	customer.HandleFunc("/summary/{org_name}", customerHandler.GetSummaryByCustomerId).Methods("GET")
	customer.HandleFunc("/top-debtors/{org_name}", customerHandler.GetTopItemsByCustomerId).Methods("GET")
	customer.HandleFunc("/top-debtors-overdue/{org_name}", customerHandler.GetTopItemsOverdueByCustomerId).Methods("GET")
	customer.HandleFunc("/blockFactors/{org_name}", customerHandler.GetCountBlockFactorsByCustomerId).Methods("GET")

	// Middleware можно добавить глобально или на подроутер
	// r.Use(middleware.Logger, middleware.Recoverer)

	return r
}
