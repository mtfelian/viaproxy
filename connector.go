package viaproxy

import "io"

// Connector is a proxy connector
type Connector interface {
	AddProxyAddr(addr string)
	GetProxyAddr() string
	RemoveProxyAddr(addr string)
	DoRequest(method string, url string, body io.Reader) (Response, error)
}
