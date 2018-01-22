package viaproxy

import "net/http"

// Response is a response for a request
type Response struct {
	*http.Response

	// ProxyAddr is a proxy address via which a request was performed
	ProxyAddr string
}
