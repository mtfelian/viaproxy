package viaproxy_test

import (
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("testing connection via proxy", func() {
	It("tests getting random proxy addr from list", func() {
		counts := map[string]int{}
		for i := 0; i < 10000; i++ { // get random addr many times (list formed in TestViaproxy())
			counts[viaSOCKS5.GetProxyAddr()]++
		}
		for _, value := range counts { // probability that at least one addr will be 0 is SMALL
			Expect(value).To(BeNumerically(">", 0))
		}
	})

	It("tests connection via proxy", func() {
		resp, err := viaSOCKS5.DoRequest(http.MethodGet, mServer.URL+"/", nil)
		Expect(err).NotTo(HaveOccurred())
		b, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(b)).To(Equal("OK"))
	})
})
