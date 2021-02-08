cli
===

## Ally Dale(vipally@gmail.com) Changelist:
- refer to [https://github.com/vipally/cmdline](https://github.com/vipally/cmdline)
- add LogicName for every flag
- add no-name flag support
- add / lead flag on windows platform
- Fix "-flag = x" or "-flag= x" or "-flag =x" cause panic bug
- add internel label <thisapp> <version> <buildtime> of app infor for descriptions
- add GetUsage() function

- v2 bug: default value changes with parsed values on slice flags
- v2 feature: support no-name flags
- v2 bug: accept multi-value input on sclice flags
- add SplitCommandLines
- avoid allocate on Flag/Command name search
- --bool=false flags
- slice flag value don't append to default values from ENV or file 

```
- add Parse function
- add no-name flags
- add Detail and Summary description
1.  Add LogicName and Required field for flag, and modify the flag define interface
2.  Add Summary and Details and Version info to usage page
3.  Add labels <thiscmd> <buildtime> <version> for getting runtime info in usage page
4.  Add interface GetUsage() string
5.  Modify Parse() logic
6.  Add no-name flag support
7.  Add "/flag" style support, named flags can lead with "/" or "-" or "--"
8.  Fix "-flag = x" or "-flag= x" or "-flag =x" cause panic bug
9.  Add synonyms support for with-name flags
10. Format usage page line head with proper num of space
```

[![GoDoc](https://godoc.org/github.com/urfave/cli?status.svg)](https://pkg.go.dev/github.com/urfave/cli/v2)
[![codebeat](https://codebeat.co/badges/0a8f30aa-f975-404b-b878-5fab3ae1cc5f)](https://codebeat.co/projects/github-com-urfave-cli)
[![Go Report Card](https://goreportcard.com/badge/urfave/cli)](https://goreportcard.com/report/urfave/cli)
[![codecov](https://codecov.io/gh/urfave/cli/branch/master/graph/badge.svg)](https://codecov.io/gh/urfave/cli)

cli is a simple, fast, and fun package for building command line apps in Go. The
goal is to enable developers to write fast and distributable command line
applications in an expressive way.

## Usage Documentation

Usage documentation exists for each major version. Don't know what version you're on? You're probably using the version from the `master` branch, which is currently `v2`.

- `v2` - [./docs/v2/manual.md](./docs/v2/manual.md)
- `v1` - [./docs/v1/manual.md](./docs/v1/manual.md)

Guides for migrating to newer versions:

- `v1-to-v2` - [./docs/migrate-v1-to-v2.md](./docs/migrate-v1-to-v2.md)

## Installation

Using this package requires a working Go environment. [See the install instructions for Go](http://golang.org/doc/install.html).

Go Modules are required when using this package. [See the go blog guide on using Go Modules](https://blog.golang.org/using-go-modules).

### Using `v2` releases

```
$ GO111MODULE=on go get github.com/urfave/cli/v2
```

```go
...
import (
  "github.com/urfave/cli/v2" // imports as package "cli"
)
...
```

### Using `v1` releases

```
$ GO111MODULE=on go get github.com/urfave/cli
```

```go
...
import (
  "github.com/urfave/cli"
)
...
```

### GOPATH

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can
be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

### Supported platforms

cli is tested against multiple versions of Go on Linux, and against the latest
released version of Go on OS X and Windows. This project uses Github Actions for
builds. To see our currently supported go versions and platforms, look at the [./.github/workflows/cli.yml](https://github.com/urfave/cli/blob/master/.github/workflows/cli.yml).
