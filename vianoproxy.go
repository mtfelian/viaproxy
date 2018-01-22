package viaproxy

import (
	"io"
	"net/http"
	"time"
)

// NewViaNoProxy returns a new proxy connector via no proxy
func NewViaNoProxy(headers http.Header) Connector { return &viaNoProxy{headers: headers} }

// viaNoProxy can made requests via no proxy
type viaNoProxy struct{ headers http.Header }

// AddProxyAddr does nothing
func (r *viaNoProxy) AddProxyAddr(_ ...string) {}

// GetProxyAddr does nothing
func (r *viaNoProxy) GetProxyAddr() string { return "" }

// RemoveProxyAddr does nothing
func (r *viaNoProxy) RemoveProxyAddr(addr string) {}

// DoRequest makes a new HTTP request via random proxy from list and returns a response
func (r *viaNoProxy) DoRequest(method, url string, body io.Reader, timeout time.Duration) (Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return Response{}, err
	}
	request.Header = r.headers

	httpResponse, err := (&http.Client{Timeout: timeout}).Do(request)
	return Response{Response: httpResponse}, err
}
