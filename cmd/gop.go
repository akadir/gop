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
	app := &cli.App{
		Name:  "gop",
		Usage: "open current git repository's remote url on browser.",
		Action: func(c *cli.Context) error {
			url := getRepositoryUrl()

			openBrowser(url)

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

func openBrowser(url string) {
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
