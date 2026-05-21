package response

import (
	"context"
	"log/slog"
	"net/http"
	"visualizationDbDebet/internal/httpx"
)

type Handler struct {
	service service
}

type service interface {
	getResponse(ctx context.Context) (*Response, error)
	getResponseWithMIP(ctx context.Context) (*Response, error)
}

func NewHandler(service service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) getResponse(w http.ResponseWriter, r *http.Request) {
	pageDTO, err := h.service.getResponse(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, pageDTO)
	slog.Info("GetResponse", responseLogAttr(pageDTO))
}

func (h *Handler) getResponseWithMIP(w http.ResponseWriter, r *http.Request) {
	pageDTO, err := h.service.getResponseWithMIP(r.Context())
	if err != nil {
		httpx.RespondError(w, err, "internal server error")
		return
	}

	httpx.RespondJSON(w, pageDTO)
	slog.Info("GetResponseWithMIP", responseLogAttr(pageDTO))
}

func responseLogAttr(response *Response) slog.Attr {
	if response == nil {
		return slog.Group("response", "present", false)
	}

	return slog.Group(
		"response",
		"present", true,
		"id", intPtrValue(response.ID),
		"count_source_org", response.CountSourceOrg,
		"count_contracts", response.CountContracts,
		"sum_contract_amount", response.SumContractAmount,
		"sum_debet_total", response.SumDebetTotal,
		"sum_debet_overdue", response.SumDebetOverdue,
	)
}

func intPtrValue(value *int) any {
	if value == nil {
		return nil
	}
	return *value
}
