package git_service

import (
	mocks "github.com/akadir/gop/mocks/gitmock"
	"github.com/akadir/gop/page"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	gitBitbucketMock *mocks.Git
)

func TestBitbucketGetPath(t *testing.T) {
	//given
	parameters := []struct {
		input    page.Page
		expected string
	}{
		{page.Pipeline, "/addon/pipelines/home"},
		{page.Mr, "/pull-requests"},
		{page.Branch, "/src/branch-name"},
		{page.Issues, "/jira"},
	}

	gitBitbucketMock = new(mocks.Git)
	gitBitbucketMock.On("GetCurrentBranchName").Return("branch-name")
	bitBucketService := NewBitbucket(gitBitbucketMock)

	for i := range parameters {
		// when
		actual := bitBucketService.GetPath(parameters[i].input)
		if actual != parameters[i].expected {
			// then
			assert.Equal(t, parameters[i].expected, actual)
		}
	}
}
