package executor

import (
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

//go:generate mockery --name=Executor --output=../../mocks/executormock
type Executor interface {
	Exec(string, ...string) []byte
}

type RealExecutor struct{}

func (r RealExecutor) Exec(command string, args ...string) []byte {
	output, err := exec.Command(command, args...).CombinedOutput()

	if err != nil {
		color.Red("%s", strings.TrimSpace(string(output)))
		color.Unset()
		os.Exit(1)
	}

	return output
}
