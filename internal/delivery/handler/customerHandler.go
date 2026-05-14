package handler

import (
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/customer"
	"visualizationBdDebet/internal/delivery/util"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service *customer.Service
}

func NewHandlerCustomer(service *customer.Service) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (h *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	customers, err := h.service.GetCustomers(ctx)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, customers)
	slog.Info("Get all customer", "customer", customers)
}

func (h *CustomerHandler) GetSummaryByCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["orgName"]

	ctx := r.Context()
	summary, err := h.service.GetSummaryByCustomerId(ctx, customerID)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, summary)
	slog.Info("Get summary by customer id", "customer", customerID)
}

func (h *CustomerHandler) GetTopItemsByCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["orgName"]

	ctx := r.Context()
	topItems, err := h.service.GetTopItemsByCustomerId(ctx, customerID)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, topItems)
	slog.Info("Get top items by customer id", "customer", customerID, "topItems", topItems)
}

func (h *CustomerHandler) GetTopItemsOverdueByCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["orgName"]

	ctx := r.Context()
	topItems, err := h.service.GetTopItemsOverdueByCustomerId(ctx, customerID)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, topItems)
	slog.Info("Get top items overdue by customer id", "customer", customerID, "topItems", topItems)
}

func (h *CustomerHandler) GetCountBlockFactorsByCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["orgName"]

	ctx := r.Context()
	factors, err := h.service.GetCountBlockFactorsByCustomerId(ctx, customerID)
	if err != nil {
		util.RespondError(w, err, "internal server error")
		return
	}

	util.RespondJSON(w, factors)
	slog.Info("Get blockFactors by customer id", "customer", customerID, "blockFactors", factors)
}
