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
	responseHandler *handler.ResponseHandler,
	customerHandler *handler.CustomerHandler,
	contractorHandler *handler.ContractorHandler,
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
	blockfactor.HandleFunc("/{id}", blockfactorHandler.GetById).Methods("GET")

	resp := r.PathPrefix("/response").Subrouter()
	resp.HandleFunc("", responseHandler.GetResponse).Methods("GET")
	resp.HandleFunc("/withMIP", responseHandler.GetResponseWithMIP).Methods("GET")

	customer := r.PathPrefix("/customer").Subrouter()
	customer.HandleFunc("", customerHandler.GetCustomers).Methods("GET")
	customer.HandleFunc("/summary/{orgName}", customerHandler.GetSummaryByCustomerId).Methods("GET")
	customer.HandleFunc("/top-debtors/{orgName}", customerHandler.GetTopItemsByCustomerId).Methods("GET")
	customer.HandleFunc("/top-debtors-overdue/{orgName}", customerHandler.GetTopItemsOverdueByCustomerId).Methods("GET")
	customer.HandleFunc("/blockFactors/{orgName}", customerHandler.GetCountBlockFactorsByCustomerId).Methods("GET")

	contractor := r.PathPrefix("/contractor").Subrouter()
	contractor.HandleFunc("/{orgName}", contractorHandler.GetContractorsWithBlockFactors).Methods("GET")
	contractor.HandleFunc("/debet/curr", contractorHandler.GetContractorsWithCurrDebet).Methods("GET")
	contractor.HandleFunc("/debet/overdue", contractorHandler.GetContractorsWithOverdueDebet).Methods("GET")

	// Middleware можно добавить глобально или на подроутер
	// r.Use(middleware.Logger, middleware.Recoverer)

	return r
}
