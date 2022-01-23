package executor

import (
	"fmt"
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
		fmt.Println(strings.TrimSpace(string(output)))
		os.Exit(1)
	}

	return output
}
