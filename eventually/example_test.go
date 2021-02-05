package example_test

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("When running", func() {

	It("should fail fast from Eventually", func() {
		defer func() {
			fmt.Fprintf(GinkgoWriter, "cleaning up in a deferred function\n")
		}()
		Eventually(func() bool {
			if _, ok := os.LookupEnv("FAIL"); ok {
				Fail("detected a terminal error, retrying is a waste of time\n")
			}
			fmt.Fprintf(GinkgoWriter, "trying again...\n")
			return false
		}, 5*time.Second, 1*time.Second).Should(BeTrue())
	})

	AfterEach(func() {
		fmt.Fprintf(GinkgoWriter, "cleaning up in AfterEach\n")
	})
})
