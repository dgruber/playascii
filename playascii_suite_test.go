package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPlayascii(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Playascii Suite")
}
