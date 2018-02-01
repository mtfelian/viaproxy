package viaproxy

import (
	"io"
	"time"
)

// Connector is a proxy connector
type Connector interface {
	AddProxyAddr(addr ...string)
	GetProxyAddr() string
	RemoveProxyAddr(addr string)
	ErrorsCount(addr string, inc bool) int
	DoRequest(method string, url string, body io.Reader, timeout time.Duration) (Response, error)
}
