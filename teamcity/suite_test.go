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
	fmt.Fprintf(GinkgoWriter, "GINKGO_REPORTERS=%s\n", rNames)

	var rs []Reporter
	if strings.Contains(rNames, "teamcity") {
		t.Log("Using TeamCity reporter")
		rs = append(rs, reporters.NewTeamCityReporter(os.Stdout))
	}
	if strings.Contains(rNames, "junit") {
		t.Log("Using JUnit reporter")
		filename := fmt.Sprintf("junit-node-%d.xml", config.GinkgoConfig.ParallelNode)
		rs = append(rs, reporters.NewJUnitReporter(filename))
	}

	fmt.Fprintf(GinkgoWriter, "GinkgoParallelNode: %d\n", GinkgoParallelNode())

	if rNames == "" || strings.Contains(rNames, "default") {
		RunSpecsWithDefaultAndCustomReporters(t, "TeamCity Reporting Experiment", rs)
	} else {
		RunSpecsWithCustomReporters(t, "TeamCity Reporting Experiment", rs)
	}

}
