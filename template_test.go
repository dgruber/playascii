package main_test

import (
	. "github.com/dgruber/playascii"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Template", func() {

	Context("HTML generation", func() {

		It("should generate the index.html", func() {
			index, err := CreateIndexTemplate()
			Ω(err).Should(BeNil())
			Ω(index).Should(ContainSubstring("demo.cast"))
		})

	})

})
