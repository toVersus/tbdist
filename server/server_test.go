package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/toversus/tbdist/store"
)

func TestGetPendingTasks(t *testing.T) {
	t.Log("getting pending tasks...")

	t.Log("should return pending tasks as JSON")

	rec := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/tasks/pending", nil)

	// The datastore is restored at the end of the test
	defer func() { ds = &store.Datastore{} }()

	ds = &mockedStore{}

	GetPendingTasks(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("KO => Got %d expected %d", rec.Code, http.StatusOK)
	}

	expect := "[{\"id\":1,\"title\":\"go to school\",\"status\":\"PENDING\"},{\"id\":2,\"title\":\"withdraw my money\",\"status\":\"PENDING\"}]"
	if result := rec.Body.String(); result != expect {
		t.Errorf("KO => Got %s expected %s", result, expect)
	}
}
