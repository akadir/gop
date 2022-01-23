package gitservice

import (
	"fmt"
	"github.com/akadir/gop/page"
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
		fmt.Println("unknown git hosting service.")
		os.Exit(1)
	}

	return gitService
}
