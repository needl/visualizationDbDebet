package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"visualizationBdDebet/internal/customer"
	"visualizationBdDebet/internal/delivery/util"
)

type Handler struct {
	service *customer.Service
}

func NewHandler(service *customer.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	customers, err := h.service.GetCustomers(ctx)
	if err != nil {
		http.Error(w, "Internal sever error", http.StatusInternalServerError)
		return
	}

	util.RespondJSON(w, customers)
	slog.Info("Get all customer: ", "customer", customers)
}

func (h *Handler) GetSummaryByCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["org_name"]
	if customerId == "" {
		http.Error(w, "Customer name is null", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	summary, err := h.service.GetSummaryByCustomerId(ctx, customerId)
	if err != nil {
		http.Error(w, "Internal sever error", http.StatusInternalServerError)
		return
	}

	/*if summary == nil {
		http.Error(w, fmt.Sprintf("Customer with name '%s' not found", customerId), http.StatusNotFound)
		return
	}*/

	util.RespondJSON(w, summary)
	slog.Info("Get summary by customer id: ", "customer", customerId)
}

func (h *Handler) GetTopItemsByCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["org_name"]
	if customerId == "" {
		http.Error(w, "Customer id is null", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	topItems, err := h.service.GetTopItemsByCustomerId(ctx, customerId)
	if err != nil {
		http.Error(w, "Internal sever error", http.StatusInternalServerError)
		return
	}

	/*if topItems == nil {
		http.Error(w, fmt.Sprintf("Customer with name '%s' not found in TopItems", customerId), http.StatusNotFound)
		return
	}*/

	util.RespondJSON(w, topItems)
	slog.Info("Get top items by customer id: ", "customer", customerId, "topItems", topItems)
}

func (h *Handler) GetTopItemsOverdueByCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["org_name"]
	if customerId == "" {
		http.Error(w, "Customer id is null", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	topItems, err := h.service.GetTopItemsOverdueByCustomerId(ctx, customerId)
	if err != nil {
		http.Error(w, "Internal sever error", http.StatusInternalServerError)
		return
	}

	/*if topItems == nil {
		http.Error(w, fmt.Sprintf("Customer with name '%s' not found in Overdue", customerId), http.StatusNotFound)
		return
	}*/

	util.RespondJSON(w, topItems)
	slog.Info("Get top items overdue by customer id: ", "customer", customerId, "topItems", topItems)
}

func (h *Handler) GetCountBlockFactorsByCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["org_name"]
	if customerId == "" {
		http.Error(w, "Customer id is null", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	factors, err := h.service.GetCountBlockFactorsByCustomerId(ctx, customerId)
	if err != nil {
		http.Error(w, "Internal sever error", http.StatusInternalServerError)
		return
	}

	if factors == nil {
		http.Error(w, fmt.Sprintf("Customer with name '%s' not found", customerId), http.StatusNotFound)
		return
	}

	util.RespondJSON(w, factors)
	slog.Info("Get blockFactors by customer id: ", "customer", customerId, "blockFactors", factors)
}
