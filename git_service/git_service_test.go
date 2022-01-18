package git_service

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func Test_Decide_With_Known_Git_Services(t *testing.T) {
	//given
	parameters := []struct {
		input    string
		expected GitService
	}{
		{"github", Github{}},
		{"gitlab", Gitlab{}},
		{"bitbucket", Bitbucket{}},
		{"git@github.com:foo/bar.git", Github{}},
		{"https://github.com/foo/bar.git", Github{}},
		{"https://gitlab.com/akadir/cv.git", Gitlab{}},
		{"git@gitlab.com:akadir/cv.git", Gitlab{}},
		{"https://bar@bitbucket.org/foo/bar.git", Bitbucket{}},
		{"git@bitbucket.org:foo/bar.git", Bitbucket{}},
	}

	for i := range parameters {
		// when
		actual := Decide(parameters[i].input)
		if actual != parameters[i].expected {
			// then
			assert.Equal(t, parameters[i].expected, actual)
		}
	}
}

func Test_Decide_With_Unknown_Git_Services(t *testing.T) {
	//given
	gitRemote := "unknownGitProvider"

	if os.Getenv("TEST_UNKNOWN_GIT_SERVICE") == "1" {
		//when
		Decide(gitRemote)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=Test_Decide_With_Unknown_Git_Services")
	cmd.Env = append(os.Environ(), "TEST_UNKNOWN_GIT_SERVICE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	assert.Failf(t, "process ran with err %v, want exit status 1", err.Error())
}
