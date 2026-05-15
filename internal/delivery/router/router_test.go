package router

import (
	"fmt"
	"sort"
	"testing"
	"visualizationDbDebet/internal/blockfactor"
	"visualizationDbDebet/internal/contract"
	"visualizationDbDebet/internal/contractor"
	"visualizationDbDebet/internal/customer"
	"visualizationDbDebet/internal/debet"
	"visualizationDbDebet/internal/object"
	"visualizationDbDebet/internal/response"

	"github.com/gorilla/mux"
)

func TestNewRouter_RegistersAllFeatureRoutes(t *testing.T) {
	r := NewRouter(
		&debet.Handler{},
		&contract.Handler{},
		&blockfactor.Handler{},
		&response.Handler{},
		&customer.Handler{},
		&contractor.Handler{},
		&object.Handler{},
	)

	got := make(map[string]struct{})
	err := r.Walk(func(route *mux.Route, _ *mux.Router, _ []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		for _, m := range methods {
			got[fmt.Sprintf("%s %s", m, path)] = struct{}{}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walk routes: %v", err)
	}

	want := []string{
		"GET /debet",
		"GET /debet/withMIP",
		"GET /contract",
		"GET /contract/{id}",
		"GET /blockFactor",
		"GET /blockFactor/{id}",
		"GET /response",
		"GET /response/withMIP",
		"GET /customer",
		"GET /customer/summary/{orgName}",
		"GET /customer/top-debtors/{orgName}",
		"GET /customer/top-debtors-overdue/{orgName}",
		"GET /customer/blockFactors/{orgName}",
		"GET /contractor/table",
		"GET /contractor/debet/curr",
		"GET /contractor/debet/overdue",
		"GET /contractor/{orgName}/debt",
		"GET /contractor/{orgName}/overdue",
		"GET /contractor/{orgName}",
		"GET /objects/search",
		"GET /objects/{sourceOrgName}",
		"GET /objects/{sourceOrgName}/{objectName}",
	}

	for _, route := range want {
		if _, ok := got[route]; !ok {
			t.Fatalf("missing route %q\nactual routes:\n%s", route, formatRoutes(got))
		}
	}
}

func formatRoutes(routes map[string]struct{}) string {
	out := make([]string, 0, len(routes))
	for r := range routes {
		out = append(out, r)
	}
	sort.Strings(out)
	return fmt.Sprint(out)
}
