package runner

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func build() (string, bool) {
	buildLog("Building Using Make file...")
	buildCommand("make", "build")

	buildLog("Building...")
	buildCommand("go", "build")

	return "", true
}

func buildCommand(command string, args string) (string, bool) {

	cmd := exec.Command(command, args, "-o", buildPath(), root())

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}
