package git

import (
	mocks "github.com/akadir/gop/mocks/executormock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	executorMock *mocks.Executor
)

func TestGitGetRepositoryUrl(t *testing.T) {
	// given
	parameters := []struct {
		mockReturn []byte
		expected   string
	}{
		{[]byte("git@github.com:foo/bar.git\n"), "https://github.com/foo/bar"},
		{[]byte("https://github.com/foo/bar.git\n"), "https://github.com/foo/bar"},
		{[]byte("git@gitlab.com:foo/bar.git\n"), "https://gitlab.com/foo/bar"},
		{[]byte("https://gitlab.com/foo/bar.git\n"), "https://gitlab.com/foo/bar"},
		{[]byte("https://foo@bitbucket.org/foo/bar.git\n"), "https://bitbucket.org/foo/bar"},
		{[]byte("git@bitbucket.org:foo/bar.git\n"), "https://bitbucket.org/foo/bar"},
	}

	for i := range parameters {
		executorMock = new(mocks.Executor)
		executorMock.On("Exec", "git", "remote", "get-url", "origin").Return(parameters[i].mockReturn)
		gitCli := NewGit(executorMock)

		// when
		repositoryUrl := gitCli.GetRepositoryUrl()
		// then
		assert.Equal(t, parameters[i].expected, repositoryUrl)
	}
}

func TestGitGetCurrentBranchName(t *testing.T) {
	// given
	executorMock = new(mocks.Executor)
	executorMock.On("Exec", "git", "rev-parse", "--abbrev-ref", "HEAD").Return([]byte("branch-name\n"))
	gitCli := NewGit(executorMock)

	// when
	branchName := gitCli.GetCurrentBranchName()
	// then
	assert.Equal(t, "branch-name", branchName)
}
