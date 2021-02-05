package example_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("When running", func() {

	It("should handle a nil pointer", func() {
		want := 1
		Eventually(func() *int {
			if _, ok := os.LookupEnv("FAIL"); ok {
				return nil
			}
			got := 1
			return &got
		}).Should(Equal(&want))
	})

})
