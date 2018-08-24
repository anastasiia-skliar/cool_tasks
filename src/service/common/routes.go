package common

import (
	"net/http"
)

// MethodHandler stores mapping for each available method-handler combinations
type MethodHandler map[string]http.Handler

// ServeHTTP checks if requested method is allowed and returns 405 if not.
// This middleware was copied from gorilla/handlers and changed according to Continuums REST API requirements
func (h MethodHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handler, ok := h[req.Method]; ok {
		handler.ServeHTTP(w, req)
	}
	SendMethodNotAllowed(w, req, "", nil)
}
