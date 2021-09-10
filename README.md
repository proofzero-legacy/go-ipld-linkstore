# go-ipld-linkstore

[![Go Reference](https://pkg.go.dev/badge/github.com/proofzero/go-ipld-linkstore.svg)](https://pkg.go.dev/github.com/proofzero/go-ipld-linkstore)
[![Go](https://img.shields.io/github/go-mod/go-version/proofzero/go-ipld-linkstore)](https://golang.org/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/proofzero/go-ipld-linkstore)](https://goreportcard.com/report/github.com/proofzero/go-ipld-linkstore)
![build](https://github.com/proofzero/go-ipld-linkstore/actions/workflows/bazel.yaml/badge.svg)
[![Slack](https://img.shields.io/badge/slack-@kubelt-FD4E83.svg)](https://kubelt.slack.com)

A small module that makes IPLD LinkSystems (newer, "prime node" architecture)
easy to use with carfiles and other, older modules that use the legacy "format
node" architecture.

# Install

At a shell within your go module:

```bash
go get github.com/proofzero/go-ipld-linkstore
```

# Build Instructions

```bash
go build
```

Bazel build:

```bash
bazel build ...
```

# Usage

Pseudo-golang for quickly and easily writing a v1 carfile full of prime nodes:

```golang
sls := NewStorageLinkSystemWithNewStorage(cidlink.DefaultLinkSystem())
cid := sls.MustStore(myLinkContext, myLinkPrototype, myPrimeNode)
car := carv1.NewSelectiveCar(context.Background(),
    sls.ReadStore, // <- special sauce block format access to prime nodes.
    []carv1.Dag{{
        // CID of the root node of the DAG to traverse.
        Root: cid.(cidlink.Link).Cid,
        // Traversal convenience selector that gives us "everything".
        Selector: everything(),
    }})
file, _ := os.Create("myV1Carfile.v1.car")
car.Write(file)
```

An attempt has been made to over-comment the code. See especially `example_test.go`.

# Testing

```bash
go test
```

```bash
bazel test ...
```
