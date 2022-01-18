package git_service

import (
	mocks "github.com/akadir/gop/mocks/gitmock"
	"github.com/akadir/gop/page"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	gitGithubMock *mocks.Git
)

func TestGithub_GetPath(t *testing.T) {
	//given
	parameters := []struct {
		input    page.Page
		expected string
	}{
		{page.Pipeline, "/actions"},
		{page.Mr, "/pulls"},
		{page.Branch, "/tree/branch-name"},
		{page.Issues, "/issues"},
	}

	gitGithubMock = new(mocks.Git)
	gitGithubMock.On("GetCurrentBranchName").Return("branch-name")
	githubService := NewGithub(gitGithubMock)

	for i := range parameters {
		// when
		actual := githubService.GetPath(parameters[i].input)
		if actual != parameters[i].expected {
			// then
			assert.Equal(t, parameters[i].expected, actual)
		}
	}
}
