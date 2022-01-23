package gitservice

import (
	"github.com/akadir/gop/cmd/executor"
	"github.com/akadir/gop/cmd/git"
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestDecideWithKnownGitServices(t *testing.T) {
	//given
	gitCli := git.NewGit(executor.RealExecutor{})

	parameters := []struct {
		input    string
		expected GitService
	}{
		{"github", NewGithub(gitCli)},
		{"gitlab", NewGitlab(gitCli)},
		{"bitbucket", NewBitbucket(gitCli)},
		{"git@github.com:foo/bar.git", NewGithub(gitCli)},
		{"https://github.com/foo/bar.git", NewGithub(gitCli)},
		{"https://gitlab.com/akadir/cv.git", NewGitlab(gitCli)},
		{"git@gitlab.com:akadir/cv.git", NewGitlab(gitCli)},
		{"https://bar@bitbucket.org/foo/bar.git", NewBitbucket(gitCli)},
		{"git@bitbucket.org:foo/bar.git", NewBitbucket(gitCli)},
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

func TestDecideWithUnknownGitServices(t *testing.T) {
	//given
	gitRemote := "unknownGitProvider"

	if os.Getenv("TEST_UNKNOWN_GIT_SERVICE") == "1" {
		//when
		Decide(gitRemote)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestDecideWithUnknownGitServices")
	cmd.Env = append(os.Environ(), "TEST_UNKNOWN_GIT_SERVICE=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	assert.Failf(t, "process ran with err %v, want exit status 1", err.Error())
}
