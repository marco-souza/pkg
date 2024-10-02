# `pkg`

CLI tool to speed-up goland development

## Installation

```bash
go install -u github.com/marco-souza/pkg

pkg
```

## Usage

```bash
pkg create <module> <folder> - scaffold a new go module
pkg encrypt <file> - encrypt a file
pkg decrypt <file> - decrypt a file
pkg -h
```

## Coming soon

```bash
pkg env set <name> <value> - set an environment variable
pkg env del <name> <value> - set an environment variable
pkg <guthub-username>/<repo> [name]- close repo as template for a new [name] project (like degit)
```
