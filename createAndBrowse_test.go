package main_test

import (
	. "github.com/dgruber/playascii"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"os"
)

var cast = `
{"version": 2, "width": 144, "height": 24, "timestamp": 1519922686, "env": {"SHELL": "/bin/bash", "TERM": "xterm-256color"}}
[0.015288, "o", "\u001b[?1034hbash-3.2$ "]
[3.033764, "o", "t"]
[3.096806, "o", "e"]
[3.289187, "o", "s"]
[3.373798, "o", "t"]
[3.49344, "o", "\r\n"]
[3.493647, "o", "bash-3.2$ "]
[3.928101, "o", "e"]
[4.120383, "o", "x"]
[4.227588, "o", "i"]
[4.335261, "o", "t"]
[4.422721, "o", "\r\n"]
[4.422802, "o", "exit\r\n"]
`

var _ = Describe("CreateAndBrowse", func() {

	Context("basic", func() {

		It("needs to create a temp dir with the html content in it", func() {
			tmpfile, err := ioutil.TempFile("", "basicCastTestFile")
			Ω(err).Should(BeNil())
			defer os.Remove(tmpfile.Name())

			tmpfile.Write([]byte(cast))
			tmpfile.Close()

			dir, path, err := CreateWebPage(tmpfile.Name())

			Ω(err).Should(BeNil())
			Ω(path).ShouldNot(Equal(""))
			Ω(dir).ShouldNot(Equal(""))
		})

	})

})
