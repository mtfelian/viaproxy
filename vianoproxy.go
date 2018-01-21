package viaproxy

import (
	"io"
	"net/http"
)

// NewViaNoProxy returns a new proxy connector via no proxy
func NewViaNoProxy() Connector { return &viaNoProxy{} }

// viaNoProxy can made requests via no proxy
type viaNoProxy struct{}

// AddProxyAddr does nothing
func (r *viaNoProxy) AddProxyAddr(_ string) {}

// GetProxyAddr does nothing
func (r *viaNoProxy) GetProxyAddr() string { return "" }

// RemoveProxyAddr does nothing
func (r *viaNoProxy) RemoveProxyAddr(addr string) {}

// DoRequest makes a new HTTP request via random proxy from list and returns a response
func (r *viaNoProxy) DoRequest(method, url string, body io.Reader) (Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return Response{}, err
	}

	httpResponse, err := http.DefaultClient.Do(request)
	return Response{Response: httpResponse}, err
}
