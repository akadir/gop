package git_service

import (
	"github.com/akadir/gop/cmd/git"
	"github.com/akadir/gop/page"
)

type Github struct{}

func (github Github) GetPath(selectedPage page.Page) string {
	var path string

	if selectedPage == page.Pipeline {
		path = "/actions"
	} else if selectedPage == page.Mr {
		path = "/pulls"
	} else if selectedPage == page.Branch {
		branchName := git.GetCurrentBranchName()
		path = "/tree/" + branchName
	} else if selectedPage == page.Issues {
		path = "/issues"
	}

	return path
}
