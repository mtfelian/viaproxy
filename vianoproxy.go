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
func (r *viaNoProxy) RemoveProxyAddr(i int) {}

// DoRequest makes a new HTTP request via random proxy from list and returns a response
func (r *viaNoProxy) DoRequest(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}
