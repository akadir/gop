package git_service

import (
	"github.com/akadir/gop/cmd/git"
	"github.com/akadir/gop/page"
)

type Bitbucket struct{}

func (bitbucket Bitbucket) GetPath(selectedPage page.Page) string {
	var path string

	if selectedPage == page.Pipeline {
		path = "/addon/pipelines/home"
	} else if selectedPage == page.Mr {
		path = "/pull-requests"
	} else if selectedPage == page.Branch {
		branchName := git.GetCurrentBranchName()
		path = "/src/" + branchName
	} else if selectedPage == page.Issues {
		path = "/jira"
	}

	return path
}
