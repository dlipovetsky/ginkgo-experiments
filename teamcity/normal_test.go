package a_test

import (
	"os"
	"time"

	"daniel.lipovetsky.me/ginkgo-experiments/util"
	. "github.com/onsi/ginkgo"
)

const (
	duration = 5 * time.Second
)

var _ = Describe("When running", func() {

	It("test A", func() {
		util.LogFor(os.Stdout, duration, "test A")
	})

	It("test B", func() {
		util.LogFor(os.Stdout, duration, "test B")
	})

})
