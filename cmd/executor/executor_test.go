package executor

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestRealExecutorExec(t *testing.T) {
	//given
	if os.Getenv("TEST_EXECUTOR_EXEC") == "1" {
		//when
		RealExecutor{}.Exec("exit", "1")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestRealExecutorExec")
	cmd.Env = append(os.Environ(), "TEST_EXECUTOR_EXEC=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	assert.Failf(t, "process ran with err %v, want exit status 1", err.Error())
}
