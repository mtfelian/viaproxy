package viaproxy

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

// regexpIP4 is an IPv4 regexp string
var regexpIP4 = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.` +
	`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.` +
	`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.` +
	`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func init() { rand.Seed(time.Now().UnixNano()) }

// NewViaSOCKS5 returns a new proxy connector via SOCKS5
func NewViaSOCKS5() Connector { return &viaSOCKS5{} }

// viaSOCKS5 can made requests via SOCKS5 proxy
type viaSOCKS5 struct {
	sync.Mutex
	Proxies []string
}

// AddProxyAddr with addr with format like IPv4:port
func (r *viaSOCKS5) AddProxyAddr(addr string) {
	if !regexp.MustCompile(fmt.Sprintf(`^%s:\d{3,5}$`, regexpIP4)).MatchString(addr) {
		return
	}
	r.Lock()
	r.Proxies = append(r.Proxies, addr)
	r.Unlock()
}

// GetProxyAddr returns random proxy addr from the list
func (r *viaSOCKS5) GetProxyAddr() string {
	r.Lock()
	defer r.Unlock()
	return r.Proxies[rand.Intn(len(r.Proxies))]
}

// RemoveProxyAddr removes i-th proxy from list
func (r *viaSOCKS5) RemoveProxyAddr(addr string) {
	r.Lock()
	defer r.Unlock()
	for i, proxyAddr := range r.Proxies {
		if proxyAddr == addr {
			r.Proxies = append(r.Proxies[:i], r.Proxies[i+1:]...)
			return
		}
	}
}

// DoRequest makes a new HTTP request via random proxy from list and returns a response
func (r *viaSOCKS5) DoRequest(method, url string, body io.Reader, timeout time.Duration) (Response, error) {
	response := Response{ProxyAddr: r.GetProxyAddr()}
	dialer, err := proxy.SOCKS5("tcp", response.ProxyAddr, nil, proxy.Direct)
	if err != nil {
		return Response{}, err
	}

	httpTransport := &http.Transport{}
	httpTransport.DialContext = func(_ context.Context, network, addr string) (net.Conn, error) {
		return dialer.Dial(network, addr)
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return response, err
	}

	httpResponse, err := (&http.Client{Transport: httpTransport, Timeout: timeout}).Do(request)
	response.Response = httpResponse
	return response, err
}
