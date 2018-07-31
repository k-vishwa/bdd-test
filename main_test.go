package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAutomataAws(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AutomataAws Suite")
}

func getPage(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	fmt.Println(string(body))
	return strings.Replace(string(body), "\n", "", 1)
}

var _ = Describe("Response", func() {
	Context("check http response", func() {
		It("get response", func() {
			responseMsg := "I got Deployed From DaaC successfully"
			host := os.Getenv("HOST")
			if host != "" {
				Expect(getPage("http://" + host + ":8080")).To(Equal(responseMsg))
			}else{
				dns := os.Getenv("DNS")
				Expect(getPage("http://" + dns)).To(Equal(responseMsg))
			}
			// Ω(getPage("http://localhost:4444")).Should(Equal("Hello World: VishwanathDevhhOps.local"))
		})

	})
})
