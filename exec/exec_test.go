package example_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	dir              = "/home/dlipovetsky/projects/konvoy-minimal-cluster"
	KonvoyExecutable = "/home/dlipovetsky/Downloads/konvoy_v1.6.0/konvoy"
	Provisioner      = "aws"
)

func Output(dir string, name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir

	return cmd.Output()
}

func TeeOutput(dir string, name string, arg ...string) (stdout, stderr []byte, err error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir
	if err := cmd.Start(); err != nil {
		return nil, nil, err
	}

	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}
	defer outPipe.Close()
	outReader := io.TeeReader(outPipe, os.Stdout)

	errPipe, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}
	defer errPipe.Close()
	errReader := io.TeeReader(errPipe, os.Stderr)

	stdout, err = ioutil.ReadAll(outReader)
	if err != nil {
		return nil, nil, err
	}

	stderr, err = ioutil.ReadAll(errReader)
	if err != nil {
		return nil, nil, err
	}

	return stdout, stderr, cmd.Wait()
}

// RunTee executes the command, allowing it to output to os.Stdout and
// os.Stderr, but returning a copy of the output as strings. This can be used,
// for example, to inspect the contents of stderr for specific error messages.
func RunTee(dir string, name string, arg ...string) (stdout string, stderr string, err error) {
	cmd := exec.Command(name, arg...)
	cmd.Dir = dir

	var outBuf bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &outBuf)

	var errBuf bytes.Buffer
	cmd.Stderr = io.MultiWriter(os.Stderr, &errBuf)

	err = cmd.Run()
	return outBuf.String(), errBuf.String(), err
}

var _ = Describe("konvoy", func() {
	It("should fail to run with error", func() {
		expectErrRegexp := ".*failed to deploy the cluster.*"
		_, stderr, err := RunTee(dir, KonvoyExecutable, "up", "--provisioner", Provisioner, "-y", "--with-checks")
		Expect(err).To(HaveOccurred(), "konvoy should exit with an error")
		Expect(string(stderr)).Should(MatchRegexp(expectErrRegexp), "the regular expression %q should have at least one match in the konvoy output", expectErrRegexp)
	})

})
