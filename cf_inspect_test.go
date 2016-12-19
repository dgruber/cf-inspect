package main_test

import (
	. "github.com/dgruber/cf-inspect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"
)

var _ = Describe("CfInspect", func() {
	Context("When not running as CF app", func() {
		It("should show the index page with null values", func() {
			req, err := http.NewRequest("GET", "/index.html", nil)
			if err != nil {
				Ω(err).To(BeNil())
			}

			r := httptest.NewRecorder()
			handler := http.HandlerFunc(IndexHandler)
			handler.ServeHTTP(r, req)

			Ω(r.Code).To(BeEquivalentTo(http.StatusOK))
			Ω(r.Body.String()).To(ContainSubstring("CF Basic Application Settings"))
		})
	})
})
