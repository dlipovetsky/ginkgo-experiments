package a_test

import (
	"os"
	"time"

	"daniel.lipovetsky.me/ginkgo-experiments/util"
	. "github.com/onsi/ginkgo"
)

const (
	duration = 1 * time.Second
)

var _ = Describe("When running", func() {

	It("should write to os.Stdout", func() {
		util.LogFor(os.Stdout, duration, "test A")
	})

	It("should write to os.Stdout, and finally fail", func() {
		util.LogFor(os.Stdout, 2*duration, "test C")
		Fail("this test is supposed to fail")
	})

	It("should write to os.Stderr", func() {
		util.LogFor(os.Stdout, 3*duration, "test B")
	})

	It("should write to GinkgoWriter", func() {
		util.LogFor(GinkgoWriter, 3*duration, "test C")
	})
})
