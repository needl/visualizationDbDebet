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
	getCustomers(ctx context.Context) ([]Customer, error)
	getSummaryByCustomerID(ctx context.Context, id string) (*Summary, error)
	getTopItemsByCustomerID(ctx context.Context, customerID string) ([]TopItem, error)
	getTopItemsOverdueByCustomerID(ctx context.Context, customerID string) ([]TopItem, error)
	getCountBlockFactorsByCustomerID(ctx context.Context, customerID string) (*BlockFactors, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.service.getCustomers(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, customers)
	slog.Info("Get all customer", "customers_count", len(customers))
}

func (h *Handler) getSummaryByCustomerID(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["orgName"]

	summary, err := h.service.getSummaryByCustomerID(r.Context(), customerID)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, summary)
	slog.Info("Get summary by customer id", "customer", customerID)
}

func (h *Handler) getTopItemsByCustomerID(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["orgName"]

	topItems, err := h.service.getTopItemsByCustomerID(r.Context(), customerID)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, topItems)
	slog.Info("Get top items by customer id", "customer", customerID, "top_items_count", len(topItems))
}

func (h *Handler) getTopItemsOverdueByCustomerID(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["orgName"]

	topItems, err := h.service.getTopItemsOverdueByCustomerID(r.Context(), customerID)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, topItems)
	slog.Info("Get top items overdue by customer id", "customer", customerID, "top_items_count", len(topItems))
}

func (h *Handler) getCountBlockFactorsByCustomerID(w http.ResponseWriter, r *http.Request) {
	customerID := mux.Vars(r)["orgName"]

	factors, err := h.service.getCountBlockFactorsByCustomerID(r.Context(), customerID)
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, factors)
	slog.Info("Get blockFactors by customer id", "customer", customerID, blockFactorsLogAttr(factors))
}

func blockFactorsLogAttr(factors *BlockFactors) slog.Attr {
	if factors == nil {
		return slog.Group("block_factors", "present", false)
	}

	return slog.Group(
		"block_factors",
		"present", true,
		"bankruptcy_count", intPtrValue(factors.BankruptcyCount),
		"liquidation_count", intPtrValue(factors.LiquidationCount),
		"unreliability_count", intPtrValue(factors.UnreliabilityCount),
		"excluded_count", intPtrValue(factors.ExcludedCount),
		"foreign_agent_count", intPtrValue(factors.ForeignAgentCount),
		"extreme_terr_count", intPtrValue(factors.ExtremeTerrCount),
		"registry_of_unscrupulous", intPtrValue(factors.RegistryOfUnscrupulous),
		"administrative_responsibility", intPtrValue(factors.AdministrativeResponsibility),
		"intention_bankrupt", intPtrValue(factors.IntentionBankrupt),
		"account_blocking_count", intPtrValue(factors.AccountBlockingCount),
		"avg_workers_list_less_than_one_count", intPtrValue(factors.AvgWorkersListLessThanOneCount),
	)
}

func intPtrValue(value *int) any {
	if value == nil {
		return nil
	}
	return *value
}
