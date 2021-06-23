
<p align="center">
<img src="logo.svg" alt="Smock Logo">
</p>

# Smock

[![GitHub release (latest by date)][release-img]][release]
[![GolangCI][golangci-lint-img]][golangci-lint]
[![Go Report Card][report-card-img]][report-card]

Instantly transform your screenshots to beautiful mockups.

## Installation

### macOS

```sh
brew install fcjr/fcjr/smock
```

### Windows

```sh
scoop bucket add fcjr https://github.com/fcjr/scoop-fcjr.git
scoop install smock
```

### Other

If you are on another platform such as linux or just rather not use a package manager, you can download a binary from Github releases and use it straight away without having to install any additional dependencies.

1) Find the latest release, download the tar.gz file for your given operating system and extract it.
2) Inside you'll find the `smock` executable which you can run directly (Note: you may need to allow execution via `chmod +x <PATH_TO_SMOCK_EXE>`).

### Development

#### Building the current commit

This project uses [goreleaser](https://github.com/goreleaser/goreleaser/).

 1) Install [go](https://golang.org/doc/install).
 2) Install goreleaser via the steps [here](https://goreleaser.com/install/).
 3) Build current commit via `goreleaser release --snapshot --skip-publish --rm-dist`.

[release-img]: https://img.shields.io/github/v/release/fcjr/smock
[release]: https://github.com/fcjr/smock/releases
[golangci-lint-img]: https://github.com/fcjr/smock/workflows/go-lint/badge.svg
[golangci-lint]: https://github.com/fcjr/smock/actions?query=workflow%3Ago-lint
[report-card-img]: https://goreportcard.com/badge/github.com/fcjr/smock
[report-card]: https://goreportcard.com/report/github.com/fcjr/smock
