package a_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

var (
	logFiles map[string]*os.File
)

func logFileForTest(name string) *os.File {
	return logFiles[name]
}

var _ = BeforeSuite(func() {
	logFiles = make(map[string]*os.File)
})

func TestExample(t *testing.T) {
	RegisterFailHandler(Fail)
	// RunSpecs(t, "Suite A")
	tcReporter := reporters.NewTeamCityReporter(os.Stdout)
	RunSpecsWithDefaultAndCustomReporters(t, "Suite A", []Reporter{tcReporter})
}
