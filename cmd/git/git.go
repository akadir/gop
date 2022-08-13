package git

import (
	"fmt"
	"github.com/akadir/gop/cmd/executor"
	"os"
	"regexp"
	"strings"
)

//go:generate mockery --name=Git --output=../../mocks/gitmock

type git struct {
	executor    executor.Executor
	remoteAlias string
}

type Git interface {
	GetRepositoryUrl() string
	GetCurrentBranchName() string
}

func NewGit(executor executor.Executor) Git {
	gitRemoteAlias := strings.TrimSpace(string(executor.Exec("git", "remote", "show")))

	if gitRemoteAlias == "" {
		fmt.Println("git remote not found in current directory. Please check git remote is set.")
		os.Exit(1)
	} else {
		gitRemoteAlias = strings.Split(gitRemoteAlias, "\n")[0]
	}

	return &git{executor: executor, remoteAlias: gitRemoteAlias}
}

func (git *git) GetRepositoryUrl() string {
	output := git.executor.Exec("git", "remote", "get-url", git.remoteAlias)

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
	output := git.executor.Exec("git", "rev-parse", "--abbrev-ref", "HEAD@{u}")

	branchName := strings.Split(strings.TrimSpace(string(output)), "\n")[0]

	branchNameParts := strings.Split(branchName, "/")

	if len(branchNameParts) > 1 {
		return strings.Join(branchNameParts[1:], "/")
	} else {
		return branchName
	}
}
