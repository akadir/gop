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
</div>

## Usage

Currently, in beta version. 

```shell
$ gop -h
NAME:
   gop - open current git repository's remote url on browser.

USAGE:
   gop [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -o value, --open value  open specific page of the repository. valid values: p pipeline w workflow mr pr
   --help, -h              show help (default: false)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details