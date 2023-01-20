package cmd

import (
	"fmt"
	"github.com/akadir/gop/autocomplete"
	"github.com/akadir/gop/cmd/executor"
	"github.com/akadir/gop/cmd/git"
	ServiceDecider "github.com/akadir/gop/gitservice"
	"github.com/akadir/gop/page"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"runtime"
)

func Run() {
	if !commandExists("git") {
		fmt.Println("git is not installed. Please install git and try again.")
		os.Exit(1)
	}

	gitCli := git.NewGit(executor.RealExecutor{})

	supportedShellTypes := make(map[string]string)
	supportedShellTypes["bash"] = autocomplete.BASH_AUTO_COMPLETE
	supportedShellTypes["zsh"] = autocomplete.ZSH_AUTO_COMPLETE
	supportedShellTypes["powershell"] = autocomplete.POWERSHELL_AUTO_COMPLETE

	app := &cli.App{
		Name:                 "gop",
		Version:              "0.8.0",
		Usage:                "gop opens current git repository's remote url on browser.",
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "https://github.com/akadir",
				Email: "",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "branch",
				Usage:    "current branch in browser.",
				Category: "open",
				Action: func(c *cli.Context) error {
					url := gitCli.GetRepositoryUrl()

					gitService := ServiceDecider.Decide(url)
					url += gitService.GetPath(page.Branch)

					openInBrowser(url)

					return nil
				},
			}, {
				Name:     "actions",
				Aliases:  []string{"pipelines"},
				Usage:    "actions/pipelines page of the repository.",
				Category: "open",
				Action: func(c *cli.Context) error {
					url := gitCli.GetRepositoryUrl()

					gitService := ServiceDecider.Decide(url)
					url += gitService.GetPath(page.Pipeline)

					openInBrowser(url)

					return nil
				},
			}, {
				Name:     "mrs",
				Aliases:  []string{"prs"},
				Usage:    "mrs/prs page of the repository.",
				Category: "open",
				Action: func(c *cli.Context) error {
					url := gitCli.GetRepositoryUrl()

					gitService := ServiceDecider.Decide(url)
					url += gitService.GetPath(page.Mr)

					openInBrowser(url)

					return nil
				},
			}, {
				Name:     "issues",
				Usage:    "issues page of the repository.",
				Category: "open",
				Action: func(c *cli.Context) error {
					url := gitCli.GetRepositoryUrl()

					gitService := ServiceDecider.Decide(url)
					url += gitService.GetPath(page.Issues)

					openInBrowser(url)

					return nil
				},
			}, {
				Name:     "settings",
				Usage:    "settings page of the repository.",
				Category: "open",
				Action: func(c *cli.Context) error {
					url := gitCli.GetRepositoryUrl()

					gitService := ServiceDecider.Decide(url)
					url += gitService.GetPath(page.Settings)

					openInBrowser(url)

					return nil
				},
			}, {
				Name:     "path",
				Usage:    "page of the given path.",
				Category: "open",
				Action: func(c *cli.Context) error {
					url := gitCli.GetRepositoryUrl()
					gitService := ServiceDecider.Decide(url)

					if c.NArg() == 0 {
						url += gitService.GetPath(page.Branch)
					} else if c.NArg() > 1 {
						fmt.Println("path command accepts only one argument.")
						os.Exit(1)
					} else {
						currentBranch := gitCli.GetCurrentBranchName()
						path := c.Args().First()
						middlePath := gitService.GetPath(page.Path)

						url = fmt.Sprintf("%s%s/%s/%s", url, middlePath, currentBranch, path)
					}

					openInBrowser(url)

					return nil
				},
			}, {
				Name:     "completion",
				Usage:    "output shell completion code for the specified shell (bash, zsh or powershell)",
				Category: "settings",
				Action: func(c *cli.Context) error {
					args := c.Args()

					shellType := args.First()

					if shellAutocompletion, ok := supportedShellTypes[shellType]; ok {
						fmt.Println(shellAutocompletion)
					} else {
						fmt.Println("unsupported shell type")
						os.Exit(1)
					}

					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			url := gitCli.GetRepositoryUrl()

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
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
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
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
