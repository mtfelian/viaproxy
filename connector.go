package viaproxy

import (
	"io"
	"net/http"
)

// Connector is a proxy connector
type Connector interface {
	AddProxyAddr(addr string)
	GetProxyAddr() string
	RemoveProxyAddr(i int)
	DoRequest(method string, url string, body io.Reader) (*http.Response, error)
}
