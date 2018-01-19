package viaproxy_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/armon/go-socks5"
	"github.com/mtfelian/viaproxy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/phayes/freeport"
)

var _ = Describe("testing connection via SOCKS5 proxy", func() {
	viaProxy := viaproxy.NewViaSOCKS5()

	// startSOCKS5Server
	startSOCKS5Server := func() {
		wg := sync.WaitGroup{}
		count := 10
		wg.Add(count)
		for i := 0; i < count; i++ {
			n, err := freeport.GetFreePort()
			Expect(err).NotTo(HaveOccurred())
			go func(port int) {
				defer GinkgoRecover()
				server, err := socks5.New(&socks5.Config{})
				Expect(err).NotTo(HaveOccurred())
				addr := fmt.Sprintf("127.0.0.1:%d", port)
				viaProxy.AddProxyAddr(addr)
				wg.Done() // can't do it after ListenAndServe(), it locks
				if err := server.ListenAndServe("tcp", addr); err != nil {
					panic(err)
				}
			}(n)
			time.Sleep(100 * time.Millisecond)
		}
		wg.Wait()
	}

	BeforeEach(func() {
		rand.Seed(time.Now().UnixNano())
		startSOCKS5Server()
	})

	It("tests getting random proxy addr from list", func() {
		counts := map[string]int{}
		for i := 0; i < 10000; i++ { // get random addr many times (list formed in TestViaproxy())
			counts[viaProxy.GetProxyAddr()]++
		}
		for _, value := range counts { // probability that at least one addr will be 0 is SMALL
			Expect(value).To(BeNumerically(">", 0))
		}
	})

	It("tests connection via proxy", func() {
		resp, err := viaProxy.DoRequest(http.MethodGet, mServer.URL+"/", nil)
		Expect(err).NotTo(HaveOccurred())
		b, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(b)).To(Equal("OK"))
	})
})
