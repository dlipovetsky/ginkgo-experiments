package a_test

import (
	"fmt"
	"os"
	"strings"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExample(t *testing.T) {
	RegisterFailHandler(Fail)

	rNames := os.Getenv("GINKGO_REPORTERS")

	var rs []Reporter
	if strings.Contains(rNames, "teamcity") {
		t.Log("Using TeamCity reporter")
		tcr := reporters.NewTeamCityReporter(os.Stdout)
		rs = append(rs, tcr)
	}
	if strings.Contains(rNames, "junit") {
		t.Log("Using JUnit reporter")
		filename := fmt.Sprintf("junit-node-%d.xml", config.GinkgoConfig.ParallelNode)
		jur := reporters.NewJUnitReporter(filename)
		rs = append(rs, jur)
	}

	if rNames == "" || strings.Contains(rNames, "default") {
		RunSpecsWithDefaultAndCustomReporters(t, "TeamCity Reporting Experiment", rs)
	} else {
		RunSpecsWithCustomReporters(t, "TeamCity Reporting Experiment", rs)
	}
}

var _ = SynchronizedAfterSuite(
	func() {
		fmt.Fprintf(os.Stderr, "GINKGO_REPORTERS=%s\n", os.Getenv("GINKGO_REPORTERS"))
		fmt.Fprintf(os.Stderr, "GinkgoParallelNode#: %d\n", GinkgoParallelNode())
		fmt.Fprintf(os.Stderr, "DefaultReporterConfig: %+v\n", config.DefaultReporterConfig)
	},
	func() {
		// no-op
	})
