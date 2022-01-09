package cmd

import (
	"fmt"
	"github.com/akadir/gop/cmd/git"
	ServiceDecider "github.com/akadir/gop/git_service"
	"github.com/akadir/gop/page"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"runtime"
)

func Run() {
	app := &cli.App{
		Name:    "gop",
		Version: "0.3.0",
		Usage:   "opens current git repository's remote url on browser.",
		Commands: []*cli.Command{
			{
				Name:  "branch",
				Usage: "opens current branch in browser.",
				Action: func(c *cli.Context) error {
					url := git.GetRepositoryUrl()

					gitService := ServiceDecider.Decide(url)
					url += gitService.GetPath(page.Branch)

					openInBrowser(url)

					return nil
				},
			},
			{
				Name:    "actions",
				Aliases: []string{"pipelines"},
				Usage:   "opens actions/pipelines page of the repository.",
				Action: func(c *cli.Context) error {
					url := git.GetRepositoryUrl()

					gitService := ServiceDecider.Decide(url)
					url += gitService.GetPath(page.Pipeline)

					openInBrowser(url)

					return nil
				},
			},
			{
				Name:    "mrs",
				Aliases: []string{"prs"},
				Usage:   "opens mrs/prs page of the repository.",
				Action: func(c *cli.Context) error {
					url := git.GetRepositoryUrl()

					gitService := ServiceDecider.Decide(url)
					url += gitService.GetPath(page.Mr)

					openInBrowser(url)

					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			url := git.GetRepositoryUrl()

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
		color.Red(err.Error())
		color.Unset()
		os.Exit(1)
	}
}
