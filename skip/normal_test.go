package b_test

import (
	"time"

	"daniel.lipovetsky.me/ginkgo-experiments/util"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("When running", func() {
	XIt("[flaky] should write to stdout", func() {
		util.LogFor(GinkgoWriter, time.Second*1, CurrentGinkgoTestDescription().FullTestText)
	})

	It("should write to stdout", func() {
		util.LogFor(GinkgoWriter, time.Second*1, CurrentGinkgoTestDescription().FullTestText)
	})
})
