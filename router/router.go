package router

import (
	"net/http"
	"regexp"
)

// Route defines regex pattern, a handler and HTTP method
type Route struct {
	Pattern    *regexp.Regexp
	Handler    http.Handler
	HTTPMethod string
}

// Router handles a list of routes
type Router struct {
	routes []*Route
}

// HandleFunc finds matched endpoint
func (r *Router) HandleFunc(pattern string, httpMethod string, f func(http.ResponseWriter, *http.Request)) {
	r.routes = append(r.routes, &Route{
		Pattern:    regexp.MustCompile(pattern + "$"),
		Handler:    http.HandlerFunc(f),
		HTTPMethod: httpMethod,
	})
}

// ServeHTTP returns while requested method and pattern is valid
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.HTTPMethod == req.Method && route.Pattern.MatchString(req.URL.Path) {
			route.Handler.ServeHTTP(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
