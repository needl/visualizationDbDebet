package contractoranalysis

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
	"time"
	"visualizationDbDebet/internal/apperr"
)

var errContractorHasNoObjects = errors.New("contractor has no objects")

var readinessNumberPattern = regexp.MustCompile(`[-+]?\d+(?:[.,]\d+)?`)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) getContractors(ctx context.Context) ([]string, error) {
	contractors, err := s.repo.findContractorNames(ctx)
	if err != nil {
		slog.Error("FindContractorNames failed", "error", err)
		return nil, err
	}

	if contractors == nil {
		return []string{}, nil
	}

	return contractors, nil
}

func (s *Service) getAnalytics(ctx context.Context, contractorName string) (*Analytics, error) {
	if contractorName == "" {
		return nil, apperr.NewInvalidArgument("contractorName is required")
	}

	summary, err := s.repo.findSummaryByContractorName(ctx, contractorName)
	if err != nil {
		slog.Error("FindSummaryByContractorName failed", "error", err)
		return nil, err
	}

	treeRows, err := s.repo.findTreeByContractorName(ctx, contractorName)
	if err != nil {
		slog.Error("FindTreeByContractorName failed", "error", err)
		return nil, err
	}

	if len(treeRows) == 0 {
		exists, err := s.repo.findContractorExists(ctx, contractorName)
		if err != nil {
			slog.Error("FindContractorExists failed", "error", err)
			return nil, err
		}
		if !exists {
			return nil, apperr.NewNotFound("contractor not found")
		}

		return nil, fmt.Errorf("%w: %s", errContractorHasNoObjects, contractorName)
	}

	customers := buildCustomerTree(treeRows)

	selected := &treeRows[0]
	details, err := s.getObjectDetails(ctx, contractorName, selected.CustomerName, selected.ObjectName)
	if err != nil {
		return nil, err
	}

	return &Analytics{
		ContractorName: contractorName,
		Summary:        *summary,
		Customers:      customers,
		SelectedObject: details,
	}, nil
}

func (s *Service) getObjectDetails(
	ctx context.Context,
	contractorName string,
	customerName string,
	objectName string,
) (*ObjectDetails, error) {
	if contractorName == "" {
		return nil, apperr.NewInvalidArgument("contractorName is required")
	}
	if customerName == "" {
		return nil, apperr.NewInvalidArgument("customerName is required")
	}
	if objectName == "" {
		return nil, apperr.NewInvalidArgument("objectName is required")
	}

	row, err := s.repo.findObjectDetails(ctx, contractorName, customerName, objectName)
	if err != nil {
		slog.Error("FindObjectDetails failed", "error", err)
		return nil, err
	}
	if row == nil {
		return nil, apperr.NewNotFound("object not found")
	}

	details := &ObjectDetails{
		Status:            statusFromMetrics(row.OverdueDebtSum, row.WorkEndDate),
		CustomerName:      customerName,
		ContractorName:    contractorName,
		ObjectName:        objectName,
		ContractSum:       row.ContractSum,
		PaidSum:           row.PaidSum,
		TDCSum:            row.TDCSum,
		RVExists:          row.RVExists,
		DebetSum:          row.DebetSum,
		OverdueDebtAmount: row.OverdueDebtSum,
		WorkStartDate:     nullableTime(row.WorkStartDate),
		WorkEndDate:       nullableTime(row.WorkEndDate),
	}

	if row.ReadinessPercent.Valid {
		value := row.ReadinessPercent.String
		details.ReadinessPercent = &value
	}

	if row.ContractSum > 0 {
		details.AcceptedPercent = row.AcceptedSum / row.ContractSum * 100
	}

	return details, nil
}

func buildCustomerTree(rows []treeRow) []CustomerNode {
	result := make([]CustomerNode, 0)
	indexByCustomer := make(map[string]int)

	for _, row := range rows {
		idx, exists := indexByCustomer[row.CustomerName]
		if !exists {
			idx = len(result)
			result = append(result, CustomerNode{
				CustomerName: row.CustomerName,
				Objects:      make([]ObjectNode, 0),
			})
			indexByCustomer[row.CustomerName] = idx
		}

		readiness := nullableReadiness(row.ReadinessPercent)

		result[idx].Objects = append(result[idx].Objects, ObjectNode{
			ObjectName:        row.ObjectName,
			ContractSum:       row.ContractSum,
			ReadinessPercent:  readiness,
			RiskLevel:         riskLevel(readiness),
			CustomerName:      row.CustomerName,
			OverdueDebtAmount: row.OverdueDebtSum,
		})
		result[idx].ObjectsCount = len(result[idx].Objects)
	}

	return result
}

func riskLevel(readiness *string) string {
	if readiness == nil {
		return "no_data"
	}

	numericReadiness, ok := parseReadinessValue(*readiness)
	if !ok {
		return "no_data"
	}

	switch {
	case numericReadiness < 30:
		return "critical"
	case numericReadiness < 70:
		return "risk"
	default:
		return "ok"
	}
}

func nullableReadiness(value sql.NullString) *string {
	if !value.Valid {
		return nil
	}

	readiness := value.String
	return &readiness
}

func parseReadinessValue(value string) (float64, bool) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return 0, false
	}

	numberMatch := readinessNumberPattern.FindString(trimmed)
	if numberMatch == "" {
		return 0, false
	}

	normalized := strings.ReplaceAll(numberMatch, ",", ".")

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return 0, false
	}

	return parsed, true
}

func statusFromMetrics(overdueDebtSum float64, workEnd sql.NullTime) string {
	if overdueDebtSum > 0 && workEnd.Valid && workEnd.Time.Before(time.Now()) {
		return "Просрочен"
	}
	return "В работе"
}

func nullableTime(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}

	t := value.Time
	return &t
}
