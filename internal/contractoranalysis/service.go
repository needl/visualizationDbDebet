package contractoranalysis

import (
	"context"
	"database/sql"
	"log/slog"
	"time"
	"visualizationDbDebet/internal/apperr"
)

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
		return nil, apperr.NewNotFound("contractor not found")
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
		Status:         statusFromMetrics(row.OverdueDebtSum, row.WorkEndDate),
		CustomerName:   customerName,
		ContractorName: contractorName,
		ObjectName:     objectName,
		ContractSum:    row.ContractSum,
		PaidSum:        row.PaidSum,
		TDCSum:         row.ContractSum,
		RVSum:          row.AcceptedSum,
		DebetSum:       row.DebetSum,
		OverdueDays:    overdueDays(row.OverdueDebtSum, row.WorkEndDate),
		WorkStartDate:  nullableTime(row.WorkStartDate),
		WorkEndDate:    nullableTime(row.WorkEndDate),
	}

	if row.ReadinessPercent.Valid {
		readiness := row.ReadinessPercent.Float64
		details.ReadinessPercent = &readiness
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

func riskLevel(readiness *float64) string {
	if readiness == nil {
		return "no_data"
	}

	switch {
	case *readiness < 30:
		return "critical"
	case *readiness < 70:
		return "risk"
	default:
		return "ok"
	}
}

func nullableReadiness(value sql.NullFloat64) *float64 {
	if !value.Valid {
		return nil
	}

	v := value.Float64
	return &v
}

func statusFromMetrics(overdueDebtSum float64, workEnd sql.NullTime) string {
	if overdueDebtSum > 0 && workEnd.Valid && workEnd.Time.Before(time.Now()) {
		return "Просрочен"
	}
	return "В работе"
}

func overdueDays(overdueDebtSum float64, workEnd sql.NullTime) int {
	if overdueDebtSum <= 0 || !workEnd.Valid {
		return 0
	}

	now := time.Now()
	if !workEnd.Time.Before(now) {
		return 0
	}

	diff := now.Sub(workEnd.Time)
	return int(diff.Hours() / 24)
}

func nullableTime(value sql.NullTime) *time.Time {
	if !value.Valid {
		return nil
	}

	t := value.Time
	return &t
}
