package router

import (
	"github.com/gorilla/mux"
	"visualizationBdDebet/internal/delivery/handler"
	"visualizationBdDebet/internal/delivery/response"
)

// NewRouter создаёт новый роутер и регистрирует все маршруты приложения.
func NewRouter(
	debetHandler *handler.DebetHandler,
	contractHandler *handler.ContractHandler,
	blockfactorHandler *handler.BlockFactorHandler,
	responseHandler *response.ResponseHandler,
) *mux.Router {
	r := mux.NewRouter()

	/*allowedPaths := map[string]bool{
		"/debet":               true,
		"/contract":            true,
		"/blockfactor":         true,
		"/response":            true,
		"/response/withoutMIP": true,
	}

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if allowedPaths[r.URL.Path] {
				next.ServeHTTP(w, r)
				return
			}
			http.NotFound(w, r)
		})
	})*/

	// Группа маршрутов для debet
	debet := r.PathPrefix("/debet").Subrouter()
	debet.HandleFunc("", debetHandler.GetAll).Methods("GET")
	debet.HandleFunc("/withMIP", debetHandler.GetAllWithMIP).Methods("GET")
	//debet.HandleFunc("/{orgName}", debetHandler.GetByOrgName).Methods("GET")

	contract := r.PathPrefix("/contract").Subrouter()
	contract.HandleFunc("", contractHandler.GetAll).Methods("GET")
	contract.HandleFunc("/{id}", contractHandler.GetById).Methods("GET")

	blockfactor := r.PathPrefix("/blockfactor").Subrouter()
	blockfactor.HandleFunc("", blockfactorHandler.GetAll).Methods("GET")
	blockfactor.HandleFunc("/{orgName}", blockfactorHandler.GetById).Methods("GET")

	resp := r.PathPrefix("/response").Subrouter()
	resp.HandleFunc("", responseHandler.GetResponse).Methods("GET")
	resp.HandleFunc("/withMIP", responseHandler.GetResponseWithMIP).Methods("GET")

	// Middleware можно добавить глобально или на подроутер
	// r.Use(middleware.Logger, middleware.Recoverer)

	return r
}
