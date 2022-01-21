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

  <!-- last commit -->
  <a href="https://github.com/akadir/muninn/commits" title="Last Commit">
     <img src="https://img.shields.io/github/last-commit/akadir/gop?style=flat" alt="last commit">
  </a>

  <!-- go report card -->
  <a href="https://goreportcard.com/report/github.com/akadir/gop" title="Go Report">
     <img src="https://goreportcard.com/badge/github.com/akadir/gop" alt="go report">
  </a>

  <!-- quality gate -->
  <a href="https://sonarcloud.io/project/overview?id=akadir_gop" title="Quality Gate">
     <img src="https://sonarcloud.io/api/project_badges/measure?project=akadir_gop&metric=alert_status" alt="quality gate">
  </a>

  <!-- code coverage -->
  <a href="https://sonarcloud.io/project/overview?id=akadir_gop" title="Code Coverage">
     <img src="https://sonarcloud.io/api/project_badges/measure?project=akadir_gop&metric=coverage" alt="code coverage">
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

Run `gop` under git repository. `gop -h` to see help.

```shell
> gop -h
NAME:
   gop - opens current git repository's remote url on browser.

USAGE:
   gop [global options] command [command options] [arguments...]

VERSION:
   0.6.1

COMMANDS:
   branch              opens current branch in browser.
   actions, pipelines  opens actions/pipelines page of the repository.
   mrs, prs            opens mrs/prs page of the repository.
   issues              opens issues page of the repository.
   settings            opens settings page of the repository.
   help, h             Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print only the version (default: false)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

<br />

<div align="center">
  Developed with ❤︎ by <a href="https://github.com/akadir">akadir</a>
</div>
