package git_service

import (
	"github.com/akadir/gop/cmd/git"
	"github.com/akadir/gop/page"
)

type Github struct {
	gitService git.Git
}

func NewGithub(gitService git.Git) GitService {
	return &Github{gitService: gitService}
}

func (github Github) GetPath(selectedPage page.Page) string {
	var path string

	if selectedPage == page.Pipeline {
		path = "/actions"
	} else if selectedPage == page.Mr {
		path = "/pulls"
	} else if selectedPage == page.Branch {
		branchName := github.gitService.GetCurrentBranchName()
		path = "/tree/" + branchName
	} else if selectedPage == page.Issues {
		path = "/issues"
	}

	return path
}
