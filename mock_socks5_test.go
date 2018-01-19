package viaproxy_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

// mockServer is a mock server
type mockServer struct{ *httptest.Server }

// imitation is a server route treating imitation
func (server *mockServer) imitation(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.RequestURI, "/") && r.Method == http.MethodGet {
		server.imitate(w, r)
		return
	}
}

// imitate is an imitation for getting the page
func (server *mockServer) imitate(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "OK") }

// newMockServer returns a pointer to new mock server
func newMockServer() *mockServer {
	var server mockServer
	return &mockServer{Server: httptest.NewServer(http.HandlerFunc((&server).imitation))}
}
