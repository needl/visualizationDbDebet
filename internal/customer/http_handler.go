package customer

import (
	"context"
	"log/slog"
	"net/http"
	"visualizationDbDebet/internal/httpx"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service
}

type service interface {
	GetCustomers(ctx context.Context) ([]Customer, error)
	GetSummaryByCustomerId(ctx context.Context, id string) (*Summary, error)
	GetTopItemsByCustomerId(ctx context.Context, customerId string) ([]TopItem, error)
	GetTopItemsOverdueByCustomerId(ctx context.Context, customerId string) ([]TopItem, error)
	GetCountBlockFactorsByCustomerId(ctx context.Context, customerId string) (*BlockFactors, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.service.GetCustomers(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, customers)
	slog.Info("Get all customer", "customer", customers)
}

func (h *Handler) GetSummaryByCustomerID(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["orgName"]

	summary, err := h.service.GetSummaryByCustomerId(r.Context(), customerID)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, summary)
	slog.Info("Get summary by customer id", "customer", customerID)
}

func (h *Handler) GetTopItemsByCustomerID(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["orgName"]

	topItems, err := h.service.GetTopItemsByCustomerId(r.Context(), customerID)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, topItems)
	slog.Info("Get top items by customer id", "customer", customerID, "topItems", topItems)
}

func (h *Handler) GetTopItemsOverdueByCustomerID(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["orgName"]

	topItems, err := h.service.GetTopItemsOverdueByCustomerId(r.Context(), customerID)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, topItems)
	slog.Info("Get top items overdue by customer id", "customer", customerID, "topItems", topItems)
}

func (h *Handler) GetCountBlockFactorsByCustomerID(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["orgName"]

	factors, err := h.service.GetCountBlockFactorsByCustomerId(r.Context(), customerID)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, factors)
	slog.Info("Get blockFactors by customer id", "customer", customerID, "blockFactors", factors)
}
