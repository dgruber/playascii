package main_test

import (
	. "github.com/dgruber/playascii"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"bytes"
)

var _ = Describe("Playascii", func() {

	Context("help", func() {

		It("should print help", func() {
			var buffer bytes.Buffer
			Help(&buffer)
			Ω(buffer.String()).Should(Equal("call with: playascii <file.cast>"))
		})

		It("should detect if the given file is correct", func() {
			err := CheckCastFile("doesNotExist")
			Ω(err).ShouldNot(BeNil())

			err = CheckCastFile("playascii.go")
			Ω(err).Should(BeNil())
		})

	})

})
