package delivery

import (
	"github.com/gorilla/mux"
	"visualizationBdDebet/internal/delivery/handler"
)

// NewRouter создаёт новый роутер и регистрирует все маршруты приложения.
func NewRouter(debetHandler *handler.DebetHandler) *mux.Router {
	r := mux.NewRouter()

	// Группа маршрутов для debet
	debet := r.PathPrefix("/debet").Subrouter()
	debet.HandleFunc("", debetHandler.GetAllDebet).Methods("GET")
	debet.HandleFunc("/{orgName}", debetHandler.GetByOrgName).Methods("GET")

	// Здесь можно добавить другие группы, например:
	// contracts := r.PathPrefix("/contracts").Subrouter()
	// contracts.HandleFunc("", contractsHandler.GetAll).Methods("GET")

	// Middleware можно добавить глобально или на подроутер
	// r.Use(middleware.Logger, middleware.Recoverer)

	return r
}
