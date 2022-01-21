package gitservice

import (
	"github.com/akadir/gop/page"
	"github.com/fatih/color"
	"os"
	"strings"
)

type GitService interface {
	GetPath(selectedPage page.Page) string
}

func Decide(url string) GitService {
	var gitService GitService

	if strings.Contains(url, "github") {
		gitService = Github{}
	} else if strings.Contains(url, "gitlab") {
		gitService = Gitlab{}
	} else if strings.Contains(url, "bitbucket") {
		gitService = Bitbucket{}
	} else {
		color.Red("unknown git hosting service.")
		color.Unset()
		os.Exit(1)
	}

	return gitService
}
