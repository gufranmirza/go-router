package router

import (
	"fmt"
	"net/http"
)

// Router serves http
type Router struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

// NewRouter creates instance of Router
func NewRouter() *Router {
	router := new(Router)
	router.handlers = make(map[string]func(http.ResponseWriter, *http.Request))
	return router
}

// ServeHTTP is called for every connection
func (s *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := s.handlers[key(r.Method, r.URL.Path)]
	if !ok {
		bad(w)
		return
	}
	f(w, r)
}

// GET sets get request
func (s *Router) GET(path string, f http.HandlerFunc) {
	s.handlers[key("GET", path)] = f
}

// POST sets post request
func (s *Router) POST(path string, f http.HandlerFunc) {
	s.handlers[key("POST", path)] = f
}

// DELETE sets delete request
func (s *Router) DELETE(path string, f http.HandlerFunc) {
	s.handlers[key("DELETE", path)] = f
}

// PUT sets put request
func (s *Router) PUT(path string, f http.HandlerFunc) {
	s.handlers[key("PUT", path)] = f
}

func key(method, path string) string {
	return fmt.Sprintf("%s:%s", method, path)
}

func bad(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"error":"not found"}`))
}
