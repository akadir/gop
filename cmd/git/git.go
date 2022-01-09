package git

import (
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

func GetRepositoryUrl() string {
	output, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()

	if err != nil {
		color.Red("%s", strings.TrimSpace(string(output)))
		color.Unset()
		os.Exit(1)
	}

	gitRemote := strings.TrimSpace(string(output))

	if strings.HasPrefix(gitRemote, "git@") {
		gitRemote = strings.Replace(gitRemote, "git@", "https://", 1)
		gitRemote = strings.Replace(gitRemote, ".com:", ".com/", 1)
		gitRemote = strings.Replace(gitRemote, ".org:", ".org/", 1)
		gitRemote = strings.Replace(gitRemote, ".git", "", 1)
	}

	return gitRemote
}

func GetCurrentBranchName() string {
	output, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").CombinedOutput()

	if err != nil {
		color.Red("%s", strings.TrimSpace(string(output)))
		color.Unset()
		os.Exit(1)
	}

	return strings.TrimSpace(string(output))
}
