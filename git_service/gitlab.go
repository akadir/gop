package git_service

import (
	"github.com/akadir/gop/cmd/git"
	"github.com/akadir/gop/page"
)

type Gitlab struct{}

func (gitlab Gitlab) GetPath(selectedPage page.Page) string {
	var path string

	if selectedPage == page.Pipeline {
		path = "/pipelines"
	} else if selectedPage == page.Mr {
		path = "/merge_requests"
	} else if selectedPage == page.Branch {
		branchName := git.GetCurrentBranchName()
		path = "/tree/" + branchName
	} else if selectedPage == page.Issues {
		path = "/issues"
	}

	return path
}
