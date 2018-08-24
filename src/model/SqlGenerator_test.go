package model

import (
	"net/http"
	"testing"
)

func TestSqlGenerator(t *testing.T) {
	var expects = "SELECT * FROM museums WHERE (1=1)"
	r, _ := http.NewRequest(http.MethodGet, "/v1/museums", nil)
	request, _, err := SQLGenerator("museums", nil, nil, r.URL.Query())
	if err != nil {
		t.Errorf("error while generating query: %s", err)
	}
	if expects != request {
		t.Error("Expected:", expects, "Was:", request)
	}
}
