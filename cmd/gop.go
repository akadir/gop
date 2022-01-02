package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Run() {
	app := &cli.App{
		Name:  "gop",
		Usage: "open current git repository's remote url on browser.",
		Commands: []*cli.Command{
			{
				Name:     "ob",
				Aliases:  []string{"open-branch"},
				Category: "open",
				Usage:    "open current branch in browser.",
				Action: func(c *cli.Context) error {
					url := getRepositoryUrl()
					branchName := getCurrentBranchName()

					url += "/tree/" + branchName

					openInBrowser(url)

					return nil
				},
			},
			{
				Name:     "op",
				Aliases:  []string{"oa", "open-pipelines", "open-actions"},
				Category: "open",
				Usage:    "open actions/pipelines page of the repository.",
				Action: func(c *cli.Context) error {
					url := getRepositoryUrl()

					url = getSpecifiedPageUrl(url, "p")

					openInBrowser(url)

					return nil
				},
			},
			{
				Name:     "omr",
				Aliases:  []string{"opr", "open-merge-requests", "open-pull-requests"},
				Category: "open",
				Usage:    "open mrs/prs page of the repository.",
				Action: func(c *cli.Context) error {
					url := getRepositoryUrl()

					url = getSpecifiedPageUrl(url, "mr")

					openInBrowser(url)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red(err.Error())
		color.Unset()
		os.Exit(1)
	}
}

func getRepositoryUrl() string {
	output, err := exec.Command("git", "remote", "get-url", "origin").Output()

	if err != nil {
		color.Red("Not a git repository: %s", err.Error())
		color.Unset()
		os.Exit(1)
	}

	gitRemote := string(output)

	if strings.HasPrefix(gitRemote, "git@") {
		gitRemote = strings.Replace(gitRemote, "git@", "https://", 1)
		gitRemote = strings.Replace(gitRemote, ".com:", ".com/", 1)
		gitRemote = strings.Replace(gitRemote, ".git", "", 1)
	}

	gitRemote = strings.TrimSpace(gitRemote)

	return gitRemote
}

func getCurrentBranchName() string {
	output, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()

	if err != nil {
		color.Red(err.Error())
		color.Unset()
		os.Exit(1)
	}

	currentBranch := string(output)

	return strings.TrimSpace(currentBranch)
}

func getSpecifiedPageUrl(url string, page string) string {
	if strings.Contains(url, "github") {
		if page == "p" || page == "pipelines" || page == "actions" || page == "a" {
			url += "/actions"
		} else if page == "mr" || page == "pr" {
			url += "/pulls"
		}
	} else if strings.Contains(url, "gitlab") {
		if page == "p" || page == "pipelines" || page == "actions" || page == "a" {
			url += "/pipelines"
		} else if page == "mr" || page == "pr" {
			url += "/merge_requests"
		}
	} else {
		color.Red("unknown git hosting service.")
		color.Unset()
		os.Exit(1)
	}

	return url
}

func openInBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
}
