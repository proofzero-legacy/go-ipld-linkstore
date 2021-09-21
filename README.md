# go-ipld-linkstore

[![Go Reference](https://pkg.go.dev/badge/github.com/proofzero/go-ipld-linkstore.svg)](https://pkg.go.dev/github.com/proofzero/go-ipld-linkstore)
[![Go](https://img.shields.io/github/go-mod/go-version/proofzero/go-ipld-linkstore)](https://golang.org/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/proofzero/go-ipld-linkstore)](https://goreportcard.com/report/github.com/proofzero/go-ipld-linkstore)
![build](https://github.com/proofzero/go-ipld-linkstore/actions/workflows/build.yaml/badge.svg)
[![matrix](https://img.shields.io/matrix/lobby:matrix.kubelt.com?label=matrix&server_fqdn=matrix.kubelt.com)](https://matrix.to/#/#lobby:matrix.kubelt.com)
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

# Contribute

We would appreciate your help to make this a useful utility. For code contributions, please send a pull request. First outlining your proposed change in an issue or discussion thread to get feedback from other developers is a good idea for anything but small changes. Other ways to contribute include:

- making a feature request
- reporting a bug
- writing a test
- adding some documentation
- providing feedback on the project direction
- reporting your experience using it

For most things we will use GitHub issues and discussions, but feel free to join the project [Matrix room](https://matrix.to/#/#lobby:matrix.kubelt.com) to chat or ask questions.
