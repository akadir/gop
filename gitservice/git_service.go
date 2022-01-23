package gitservice

import (
	"fmt"
	"github.com/akadir/gop/cmd/executor"
	"github.com/akadir/gop/cmd/git"
	"github.com/akadir/gop/page"
	"os"
	"strings"
)

type GitService interface {
	GetPath(selectedPage page.Page) string
}

func Decide(url string) GitService {
	var gitService GitService
	var gitCli = git.NewGit(executor.RealExecutor{})

	if strings.Contains(url, "github") {
		gitService = NewGithub(gitCli)
	} else if strings.Contains(url, "gitlab") {
		gitService = NewGitlab(gitCli)
	} else if strings.Contains(url, "bitbucket") {
		gitService = NewBitbucket(gitCli)
	} else {
		fmt.Println("unknown git hosting service.")
		os.Exit(1)
	}

	return gitService
}
