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
		Name:    "gop",
		Version: "0.2.2",
		Usage:   "open current git repository's remote url on browser.",
		Commands: []*cli.Command{
			{
				Name:  "current-branch",
				Usage: "open current branch in browser.",
				Action: func(c *cli.Context) error {
					url := getRepositoryUrl()

					url = getSpecifiedPageUrl(url, "current-branch")

					openInBrowser(url)

					return nil
				},
			},
			{
				Name:    "actions",
				Aliases: []string{"pipelines"},
				Usage:   "open actions/pipelines page of the repository.",
				Action: func(c *cli.Context) error {
					url := getRepositoryUrl()

					url = getSpecifiedPageUrl(url, "p")

					openInBrowser(url)

					return nil
				},
			},
			{
				Name:    "mrs",
				Aliases: []string{"prs"},
				Usage:   "open mrs/prs page of the repository.",
				Action: func(c *cli.Context) error {
					url := getRepositoryUrl()

					url = getSpecifiedPageUrl(url, "mr")

					openInBrowser(url)

					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			url := getRepositoryUrl()

			openInBrowser(url)

			return nil
		},
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red(err.Error())
		color.Unset()
		os.Exit(1)
	}
}

func getRepositoryUrl() string {
	output, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()

	if err != nil {
		color.Red("%s", strings.TrimSpace(string(output)))
		color.Unset()
		os.Exit(1)
	}

	gitRemote := string(output)

	if strings.HasPrefix(gitRemote, "git@") {
		gitRemote = strings.Replace(gitRemote, "git@", "https://", 1)
		gitRemote = strings.Replace(gitRemote, ".com:", ".com/", 1)
		gitRemote = strings.Replace(gitRemote, ".org:", ".org/", 1)
		gitRemote = strings.Replace(gitRemote, ".git", "", 1)
	}

	gitRemote = strings.TrimSpace(gitRemote)

	return gitRemote
}

func getCurrentBranchName() string {
	output, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").CombinedOutput()

	if err != nil {
		color.Red("%s", strings.TrimSpace(string(output)))
		color.Unset()
		os.Exit(1)
	}

	currentBranch := string(output)

	return strings.TrimSpace(currentBranch)
}

func getSpecifiedPageUrl(url string, page string) string {
	if strings.Contains(url, "github") {
		if page == "p" {
			url += "/actions"
		} else if page == "mr" {
			url += "/pulls"
		} else if page == "current-branch" {
			branchName := getCurrentBranchName()
			url += "/tree/" + branchName
		}
	} else if strings.Contains(url, "gitlab") {
		if page == "p" {
			url += "/pipelines"
		} else if page == "mr" {
			url += "/merge_requests"
		} else if page == "current-branch" {
			branchName := getCurrentBranchName()
			url += "/tree/" + branchName
		}
	} else if strings.Contains(url, "bitbucket") {
		if page == "p" {
			url += "/addon/pipelines/home"
		} else if page == "mr" {
			url += "/pull-requests"
		} else if page == "current-branch" {
			branchName := getCurrentBranchName()
			url += "/src/" + branchName
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
