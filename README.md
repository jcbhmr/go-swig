# SWIG Go distribution

🥤 SWIG distributed as a `go install`-able module

## Installation

```sh
go install github.com/jcbhmr/go-swig/cmd/swig@latest
```

ℹ The `swig` binary will extract, build, and compile the actual native SWIG binary the first time it is run. This may take a few minutes. You can preempt this during the install process by running `swig -version` to trigger the extraction and compilation process without doing anything meaningful.

There's a second less-used binary in the SWIG project: `ccache-swig`. This binary is **not available until you run `swig` at least once**. I wrapper for it can also be installed with `go install github.com/jcbhmr/go-swig/cmd/ccache-swig@latest`, in which case the regular `swig` binary **will not be available until you run `ccache-swig` at least once**.

<details><summary>Uninstalling</summary>

```sh
rm -rf "$(go env GOPATH)"/bin/{.go-swig,swig,ccache-swig}
```

</details>

## Usage

You can use the `swig` command as though it were the original SWIG command! 🚀

```sh
swig -help
```

## How?

1. You run `go install github.com/jcbhmr/go-swig/cmd/swig@latest` which fetches, compiles, and installs the `swig` Go wrapper binary from an embedded `.tar.gz` `go:embed`. Use `-tags noembed` to get a smaller binary that downloads the `.tar.gz` from the internet.
2. You run the Go-based `swig` wrapper. This triggers the following:
    1. If `-tags noembed`, download the SWIG source tarball from the internet.
    2. Extract the tarball into the `.go-swig` folder next to the `swig` Go-based wrapper binary's location.
    3. Run the SWIG build process if on Linux or macOS. On Windows there's already a prebuilt static binary.
3. Replace the Go wrapper binary with a symlink to the actual `swig` binary in the `.go-swig` folder.
4. Now that there is the actual `swig` native binary in the `.go-swig` folder, replace the current process with the actual `swig` binary and pass along the arguments & environment. On Windows this spawns a transparent subprocess.
