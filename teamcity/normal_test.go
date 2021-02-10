package a_test

import (
	"os"
	"sync"
	"time"

	"daniel.lipovetsky.me/ginkgo-experiments/util"
	. "github.com/onsi/ginkgo"
)

const (
	duration = 1 * time.Second
)

var _ = Describe("When running", func() {

	It("should write to os.Stdout", func() {
		util.LogFor(os.Stdout, 3*duration, "should write to os.Stdout")
	})

	It("should write to os.Stderr", func() {
		util.LogFor(os.Stderr, 2*duration, "should write to os.Stderr")
	})

	It("should write to os.Stdout, os.Stderr, and finally fail", func() {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			util.LogFor(os.Stdout, 3*duration, "should write to os.Stdout")
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			go util.LogFor(os.Stderr, 3*duration, "should write to os.Stderr")
			wg.Done()
		}()
		wg.Wait()
		Fail("this test is supposed to fail")
	})

	It("should write to GinkgoWriter", func() {
		util.LogFor(GinkgoWriter, 3*duration, "should write to GinkgoWriter")
	})
})
