package gitService

import (
	mocks "github.com/akadir/gop/mocks/gitmock"
	"github.com/akadir/gop/page"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	gitGitlabMock *mocks.Git
)

func TestGitlabGetPath(t *testing.T) {
	//given
	parameters := []struct {
		input    page.Page
		expected string
	}{
		{page.Pipeline, "/pipelines"},
		{page.Mr, "/merge_requests"},
		{page.Branch, "/tree/branch-name"},
		{page.Issues, "/issues"},
		{page.Settings, "/edit"},
	}

	gitGitlabMock = new(mocks.Git)
	gitGitlabMock.On("GetCurrentBranchName").Return("branch-name")
	githubService := NewGitlab(gitGitlabMock)

	for i := range parameters {
		// when
		actual := githubService.GetPath(parameters[i].input)
		// then
		assert.Equal(t, parameters[i].expected, actual)
	}
}
