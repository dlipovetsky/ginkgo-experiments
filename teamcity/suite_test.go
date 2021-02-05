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

	var rs []Reporter
	if strings.Contains(os.Getenv("GINKGO_REPORTERS"), "teamcity") {
		t.Log("Using TeamCity reporter")
		rs = append(rs, reporters.NewTeamCityReporter(GinkgoWriter))
	}
	if strings.Contains(os.Getenv("GINKGO_REPORTERS"), "junit") {
		t.Log("Using JUnit reporter")
		filename := fmt.Sprintf("junit-node-%d.xml", config.GinkgoConfig.ParallelNode)
		rs = append(rs, reporters.NewJUnitReporter(filename))
	}

	RunSpecsWithDefaultAndCustomReporters(t, "TeamCity Reporting Experiment", rs)
}
