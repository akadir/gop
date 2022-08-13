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
		mockReturnOrigin        []byte
		mockReturnRepositoryUrl []byte
		expected                string
	}{
		{[]byte("origin"), []byte("git@github.com:foo/bar.git\n"), "https://github.com/foo/bar"},
		{[]byte("origin"), []byte("https://github.com/foo/bar.git\n"), "https://github.com/foo/bar"},
		{[]byte("origin"), []byte("git@gitlab.com:foo/bar.git\n"), "https://gitlab.com/foo/bar"},
		{[]byte("origin"), []byte("https://gitlab.com/foo/bar.git\n"), "https://gitlab.com/foo/bar"},
		{[]byte("origin"), []byte("https://foo@bitbucket.org/foo/bar.git\n"), "https://bitbucket.org/foo/bar"},
		{[]byte("origin"), []byte("git@bitbucket.org:foo/bar.git\n"), "https://bitbucket.org/foo/bar"},
	}

	for i := range parameters {
		executorMock = new(mocks.Executor)

		executorMock.On("Exec", "git", "remote", "show").Return(parameters[i].mockReturnOrigin)
		executorMock.On("Exec", "git", "remote", "get-url", string(parameters[i].mockReturnOrigin)).Return(parameters[i].mockReturnRepositoryUrl)

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
	executorMock.On("Exec", "git", "remote", "show").Return([]byte("origin"))
	executorMock.On("Exec", "git", "rev-parse", "--abbrev-ref", "HEAD@{u}").Return([]byte("branch-name\n"))
	gitCli := NewGit(executorMock)

	// when
	branchName := gitCli.GetCurrentBranchName()
	// then
	assert.Equal(t, "branch-name", branchName)
}

func TestGitGetCurrentBranchNameWithRemotePrefix(t *testing.T) {
	// given
	executorMock = new(mocks.Executor)
	executorMock.On("Exec", "git", "remote", "show").Return([]byte("origin"))
	executorMock.On("Exec", "git", "rev-parse", "--abbrev-ref", "HEAD@{u}").Return([]byte("origin/branch-name\n"))
	gitCli := NewGit(executorMock)

	// when
	branchName := gitCli.GetCurrentBranchName()
	// then
	assert.Equal(t, "branch-name", branchName)
}
