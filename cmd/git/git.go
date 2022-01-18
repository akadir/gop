package git

import (
	"github.com/akadir/gop/cmd/executor"
	"regexp"
	"strings"
)

//go:generate mockery --name=Git --output=../../mocks/gitmock
type git struct {
	executor executor.Executor
}

type Git interface {
	GetRepositoryUrl() string
	GetCurrentBranchName() string
}

func NewGit(executor executor.Executor) Git {
	return &git{executor: executor}
}

func (git *git) GetRepositoryUrl() string {
	output := git.executor.Exec("git", "remote", "get-url", "origin")

	gitRemote := strings.TrimSpace(string(output))

	if strings.HasPrefix(gitRemote, "git@") {
		gitRemote = strings.Replace(gitRemote, ":", "/", 1)
		gitRemote = strings.Replace(gitRemote, "git@", "https://", 1)
	}

	bitbucketPrefix := regexp.MustCompile(`\w*@`)
	gitRemote = bitbucketPrefix.ReplaceAllString(gitRemote, "")

	gitRemote = strings.Replace(gitRemote, ".git", "", 1)

	return gitRemote
}

func (git *git) GetCurrentBranchName() string {
	output := git.executor.Exec("git", "rev-parse", "--abbrev-ref", "HEAD")

	return strings.TrimSpace(string(output))
}
