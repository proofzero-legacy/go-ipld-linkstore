# go-ipld-linkstore

[![Go Reference](https://pkg.go.dev/badge/kubelt.com/go-ipld-linkstore.svg)](https://pkg.go.dev/kubelt.com/go-ipld-linkstore)
[![Go](https://img.shields.io/github/go-mod/go-version/proofzero/go-ipld-linkstore)](https://golang.org/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/proofzero/go-ipld-linkstore)](https://goreportcard.com/report/github.com/proofzero/go-ipld-linkstore)
[![build](https://github.com/proofzero/kmdr/actions/workflows/bazel.yaml/badge.svg)]()
[![platforms](https://img.shields.io/badge/platforms-linux|windows|macos-inactive.svg)]()
[![Slack](https://img.shields.io/badge/slack-@kubelt-FD4E83.svg)](https://kubelt.slack.com)

A small module that makes IPLD LinkSystems (newer, "prime node" architecture)
easy to use with carfiles and other, older modules that use the legacy "format
node" architecture.

# Install

At a shell within your go module:

```bash
$ go get github.com/proofzero/go-ipld-linkstore
```

# Build Instructions

```bash
$ go build
```

Bazel build coming soon.

# Usage

An attempt has been made to over-comment the code. See also the tests:

```bash
$ go test
```
