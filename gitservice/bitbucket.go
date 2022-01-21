package gitservice

import (
	"github.com/akadir/gop/cmd/git"
	"github.com/akadir/gop/page"
)

type Bitbucket struct {
	gitService git.Git
}

func NewBitbucket(gitService git.Git) GitService {
	return &Bitbucket{gitService: gitService}
}

func (bitbucket Bitbucket) GetPath(selectedPage page.Page) string {
	var path string

	if selectedPage == page.Pipeline {
		path = "/addon/pipelines/home"
	} else if selectedPage == page.Mr {
		path = "/pull-requests"
	} else if selectedPage == page.Branch {
		branchName := bitbucket.gitService.GetCurrentBranchName()
		path = "/src/" + branchName
	} else if selectedPage == page.Issues {
		path = "/jira"
	} else if selectedPage == page.Settings {
		path = "/admin"
	}

	return path
}
