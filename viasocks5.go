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
func NewViaSOCKS5(headers http.Header) Connector { return &viaSOCKS5{headers: headers} }

// viaSOCKS5 can made requests via SOCKS5 proxy
type viaSOCKS5 struct {
	sync.Mutex
	headers http.Header
	proxies []string
}

// AddProxyAddr with addr with format like IPv4:port
func (r *viaSOCKS5) AddProxyAddr(addr ...string) {
	for _, address := range addr {
		if !regexp.MustCompile(fmt.Sprintf(`^%s:\d{3,5}$`, regexpIP4)).MatchString(address) {
			continue
		}
		r.Lock()
		r.proxies = append(r.proxies, address)
		r.Unlock()
	}
}

// GetProxyAddr returns random proxy addr from the list
func (r *viaSOCKS5) GetProxyAddr() string {
	r.Lock()
	defer r.Unlock()
	return r.proxies[rand.Intn(len(r.proxies))]
}

// RemoveProxyAddr removes i-th proxy from list
func (r *viaSOCKS5) RemoveProxyAddr(addr string) {
	r.Lock()
	defer r.Unlock()
	for i, proxyAddr := range r.proxies {
		if proxyAddr == addr {
			r.proxies = append(r.proxies[:i], r.proxies[i+1:]...)
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
	request.Header = r.headers

	httpResponse, err := (&http.Client{Transport: httpTransport, Timeout: timeout}).Do(request)
	response.Response = httpResponse
	return response, err
}
