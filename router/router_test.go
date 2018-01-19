package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var routeTests = []struct {
	name       string
	routMethod string
	routPath   string
	reqMethod  string
	reqURL     string
	expect     int
}{
	{
		name:       "should response with a status 200 OK when a route and method match",
		routPath:   "/tasks",
		routMethod: http.MethodGet,
		reqURL:     "/tasks",
		reqMethod:  http.MethodGet,
		expect:     http.StatusOK,
	},
	{
		name:       "should response with a status 404 Not Found when HTTP method is different",
		routPath:   "/tasks",
		routMethod: http.MethodGet,
		reqURL:     "/tasks",
		reqMethod:  http.MethodPost,
		expect:     http.StatusNotFound,
	},
	{
		name:       "should response with a status 200 OK when a route match regex and method",
		routPath:   `/tasks/\d`,
		routMethod: http.MethodGet,
		reqURL:     "/tasks/1",
		reqMethod:  http.MethodGet,
		expect:     http.StatusOK,
	},
	{
		name:       "should response with a statu 404 Not Found when route could not be found",
		routPath:   `/tasks\d`,
		routMethod: http.MethodPost,
		reqURL:     "/tasks/a",
		reqMethod:  http.MethodPost,
		expect:     http.StatusNotFound,
	},
}

func TestRoute(t *testing.T) {
	t.Log("routing...")

	for _, testcase := range routeTests {
		rec := httptest.NewRecorder()
		req, _ := http.NewRequest(testcase.reqMethod, testcase.reqURL, nil)

		r := Route{}

		r.HandleFunc(testcase.routPath, testcase.routMethod, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.ServerHTTP(rec, req)

		if rec.Code != testcase.expect {
			t.Errorf("KO => Get %d expected %d", rec.Code, testcase.expect)
		}
	}
}
