<h1 align="center">gop</h1>

<div align="center">
  simple cli app to open current git repository's remote url in the default browser.
</div>

<br>

<div align="center">
  <!-- go version -->
  <a href="https://github.com/akadir/gop" title="go version">
    <img src="https://img.shields.io/badge/go-1.17-black.svg" alt="go version"/>
  </a>

  <!-- CI -->
  <a href="https://github.com/akadir/gop/actions" title="build status">
    <img src="https://github.com/akadir/gop/actions/workflows/build.yml/badge.svg" alt="CI"/>
  </a>

  <!-- License -->
  <a href="https://img.shields.io/badge/License-MIT-blue.svg">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg"
      alt="License" />
  </a>
</div>

## Installation

```shell
brew install akadir/gop/gop
```
or
```shell
brew tap akadir/gop && brew install gop
```

## Usage

Currently, in **beta** version.

Run `gop` under git repository. `gop -h` to show help.

```shell
> gop -h
NAME:
   gop - open current git repository's remote url on browser.

USAGE:
   gop [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command
   open:
     ob, open-branch                                    open current branch in browser.
     op, oa, open-pipelines, open-actions               open actions/pipelines page of the repository.
     omr, opr, open-merge-requests, open-pull-requests  open mrs/prs page of the repository.

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details