package viaproxy

import (
	"io"
	"time"
)

// Connector is a proxy connector
type Connector interface {
	AddProxyAddr(addr string)
	GetProxyAddr() string
	RemoveProxyAddr(addr string)
	DoRequest(method string, url string, body io.Reader, timeout time.Duration) (Response, error)
}
