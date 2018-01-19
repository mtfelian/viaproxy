package viaproxy

import (
	"io"
	"net/http"
)

// ProxyConnector is a proxy connector
type ProxyConnector interface {
	AddProxyAddr(string)
	GetProxyAddr() string
	RemoveProxyAddr(int)
	DoRequest(string, string, io.Reader) (*http.Response, error)
}
