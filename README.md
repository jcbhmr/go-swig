# SWIG Go distribution

🥤 SWIG distributed as a `go install`-able module

## Installation

```sh
go install github.com/jcbhmr/go-swig/cmd/...@latest
```

```sh
go get github.com/jcbhmr/go-swig
```

## Usage

You can use the `swig` command as though it were the original SWIG command! 🚀

```sh
swig -help
ccache-swig -help
```

```go
//go:generate go run github.com/jcbhmr/go-swig/cmd/swig ...
```
