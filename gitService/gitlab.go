package gitService

import (
	"github.com/akadir/gop/cmd/git"
	"github.com/akadir/gop/page"
)

type Gitlab struct {
	gitService git.Git
}

func NewGitlab(gitService git.Git) GitService {
	return &Gitlab{gitService: gitService}
}
func (gitlab Gitlab) GetPath(selectedPage page.Page) string {
	var path string

	if selectedPage == page.Pipeline {
		path = "/pipelines"
	} else if selectedPage == page.Mr {
		path = "/merge_requests"
	} else if selectedPage == page.Branch {
		branchName := gitlab.gitService.GetCurrentBranchName()
		path = "/tree/" + branchName
	} else if selectedPage == page.Issues {
		path = "/issues"
	} else if selectedPage == page.Settings {
		path = "/edit"
	}

	return path
}
