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
			host := os.Getenv("HOST")
			Expect(getPage("http://" + host + ":8081")).To(Equal("Hello World-Update"))
			// Ω(getPage("http://localhost:4444")).Should(Equal("Hello World: VishwanathDevhhOps.local"))
		})

	})
})
