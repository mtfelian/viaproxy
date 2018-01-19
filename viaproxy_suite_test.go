package viaproxy_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/armon/go-socks5"
	"github.com/mtfelian/viaproxy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	mServer   *mockServer
	viaSOCKS5 = viaproxy.NewViaSOCKS5()
)

// startSOCKS5Server
func startSOCKS5Server() {
	wg := sync.WaitGroup{}
	count, startPort := 10, 8030
	wg.Add(count)
	for i := startPort; i < startPort+count; i++ {
		go func(port int) {
			defer GinkgoRecover()
			server, err := socks5.New(&socks5.Config{})
			if err != nil {
				panic(err)
			}
			addr := fmt.Sprintf("127.0.0.1:%d", port)
			viaSOCKS5.AddProxyAddr(addr)
			wg.Done() // can't do it after ListenAndServe(), it locks
			if err := server.ListenAndServe("tcp", addr); err != nil {
				panic(err)
			}
		}(i)
	}
	wg.Wait()
}

// TestViaproxy launches test suite
func TestViaproxy(t *testing.T) {
	// init mock server
	mServer = newMockServer()
	defer mServer.Close()

	rand.Seed(time.Now().UnixNano())
	startSOCKS5Server()
	time.Sleep(time.Second)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Viaproxy Suite")
}
