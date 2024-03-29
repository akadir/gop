package gitservice

import (
	mocks "github.com/akadir/gop/mocks/gitmock"
	"github.com/akadir/gop/page"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	gitGithubMock *mocks.Git
)

func TestGithubGetPath(t *testing.T) {
	//given
	parameters := []struct {
		input    page.Page
		expected string
	}{
		{page.Pipeline, "/actions"},
		{page.Mr, "/pulls"},
		{page.Branch, "/tree/branch-name"},
		{page.Issues, "/issues"},
		{page.Settings, "/settings"},
		{page.Path, "/blob"},
	}

	gitGithubMock = new(mocks.Git)
	gitGithubMock.On("GetCurrentBranchName").Return("branch-name")
	githubService := NewGithub(gitGithubMock)

	for i := range parameters {
		// when
		actual := githubService.GetPath(parameters[i].input)
		// then
		assert.Equal(t, parameters[i].expected, actual)
	}
}
