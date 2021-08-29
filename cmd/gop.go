package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Run() {
	validPages := []string{"p", "pipeline", "w", "workflow", "mr", "pr"}

	app := &cli.App{
		Name:  "gop",
		Usage: "open current git repository's remote url on browser.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "o",
				Aliases: []string{"open"},
				Usage:   "open specific page of the repository. valid values: " + strings.Join(validPages, " "),
			},
		},
		Action: func(c *cli.Context) error {
			url := getRepositoryUrl()

			specifiedPage := c.String("o")

			if specifiedPage != "" {
				ok := false

				for _, valid := range validPages {
					if specifiedPage == valid {
						ok = true
					}
				}

				if !ok {
					return fmt.Errorf("page must be one of %v", validPages)
				}

				url = getSpecifiedPageUrl(url, specifiedPage)
			}
			fmt.Printf(url)
			openInBrowser(url)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getRepositoryUrl() string {
	output, err := exec.Command("git", "remote", "get-url", "origin").Output()

	if err != nil {
		log.Fatal("Not a git repository: ", err.Error())
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

func getSpecifiedPageUrl(url string, page string) string {
	if strings.Contains(url, "github") {
		if page == "p" || page == "pipeline" || page == "workflow" || page == "w" {
			url += "/actions"
		} else if page == "mr" || page == "pr" {
			url += "/pulls"
		}
	} else if strings.Contains(url, "gitlab") {
		if page == "p" || page == "pipeline" || page == "workflow" || page == "w" {
			url += "/pipelines"
		} else if page == "mr" || page == "pr" {
			url += "/merge_requests"
		}
	} else {
		fmt.Printf("unknown git hosting service.")
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
		log.Fatal(err)
	}
}
