package a_test

import (
	"fmt"
	"os"
	"time"

	"daniel.lipovetsky.me/ginkgo-experiments/util"
	. "github.com/onsi/ginkgo"
)

const (
	duration = 2 * time.Second
)

var _ = Describe("When running", func() {
	BeforeEach(func() {
		logFile, err := os.Create(fmt.Sprintf("%s.log", CurrentGinkgoTestDescription().FullTestText))
		if err != nil {
			Fail(fmt.Sprintf("could not configure log for test: %s", err))
		}
		logFiles[CurrentGinkgoTestDescription().FullTestText] = logFile

		util.LogFor(logFile, duration, fmt.Sprintf("BeforeEach for %s:", CurrentGinkgoTestDescription().FullTestText))
	})

	AfterEach(func() {
		util.LogFor(logFileForTest(CurrentGinkgoTestDescription().FullTestText), duration, fmt.Sprintf("AfterEach for %s:", CurrentGinkgoTestDescription().FullTestText))
	})

	It("should log to file", func() {
		util.LogFor(logFileForTest(CurrentGinkgoTestDescription().FullTestText), duration, CurrentGinkgoTestDescription().FullTestText)
	})

	It("should _also_ log to file", func() {
		util.LogFor(logFileForTest(CurrentGinkgoTestDescription().FullTestText), 8*duration, CurrentGinkgoTestDescription().FullTestText)
	})

	// control (do not change output destination)
	It("should log to GinkgoWriter", func() {
		util.LogFor(GinkgoWriter, duration, CurrentGinkgoTestDescription().FullTestText)
	})

	// control (do not change output destination)
	It("should log to os.Stdout", func() {
		util.LogFor(os.Stdout, duration, CurrentGinkgoTestDescription().FullTestText)
	})

	// control (do not change output destination)
	It("should log to os.Stderr", func() {
		util.LogFor(os.Stderr, duration, CurrentGinkgoTestDescription().FullTestText)
	})

})
