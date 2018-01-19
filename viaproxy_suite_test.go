package viaproxy_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var mServer *mockServer

// TestViaproxy launches test suite
func TestViaproxy(t *testing.T) {
	// init mock server
	mServer = newMockServer()
	defer mServer.Close()

	RegisterFailHandler(Fail)
	RunSpecs(t, "Viaproxy Suite")
}
