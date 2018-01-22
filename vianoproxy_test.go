package viaproxy_test

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/mtfelian/viaproxy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("testing connection via no proxy", func() {
	viaProxy := viaproxy.NewViaNoProxy(nil)

	It("tests connection without proxy", func() {
		resp, err := viaProxy.DoRequest(http.MethodGet, mServer.URL+"/", nil, time.Second)
		Expect(err).NotTo(HaveOccurred())
		b, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(b)).To(Equal("OK"))
	})
})
