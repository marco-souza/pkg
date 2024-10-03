# `pkg`

CLI tool to speed-up goland development

![GitHub](https://img.shields.io/badge/github-%23121011.svg?style=for-the-badge&logo=github&logoColor=white)
[![Go Reference](https://pkg.go.dev/badge/github.com/marco-souza/pkg.svg)](https://pkg.go.dev/github.com/marco-souza/pkg)
[![Go Report Card](https://goreportcard.com/badge/github.com/marco-souza/pkg)](https://goreportcard.com/report/github.com/marco-souza/pkg)

## Features

- scaffold a new go module
- encrypt and decrypt files
- ...

## Installation

```bash
go install -u github.com/marco-souza/pkg@latest

pkg
```

## Usage

```bash
pkg -h

pkg create <module> <folder> - scaffold a new go module
pkg encrypt <file> - encrypt a file
pkg decrypt <file> - decrypt a file
```

## Coming soon

```bash
pkg env set <name> <value> - set an environment variable
pkg env del <name> <value> - set an environment variable
pkg <guthub-username>/<repo> [name]- close repo as template for a new [name] project (like degit)
```

## References

- [cobra](https://github.com/spf13/cobra)
- [cobra-cli](https://github.com/spf13/cobra-cli)

