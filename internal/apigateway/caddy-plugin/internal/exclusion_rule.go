package internal


import "net/http"

const AllMethod = "*"

var HttpMethods = []string{AllMethod, http.MethodPost, http.MethodDelete,
	http.MethodPatch, http.MethodPut, http.MethodGet, http.MethodOptions, http.MethodConnect}

// Path exclusion rule
type ExclusionRule struct {
	Method string
	Path   string
}