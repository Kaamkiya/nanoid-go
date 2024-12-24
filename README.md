<center>
# nanoid

A small, fast library for generating cryptographically secure, URL safe IDs.

![Go Report Card](https://goreportcard.com/badge/github.com/Kaamkiya/nanoid-go)
![License: Unlicense](https://img.shields.io/badge/License-Unlicense-blue.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/Kaamkiya/nanoid-go.svg)](https://pkg.go.dev/github.com/Kaamkiya/nanoid-go)
</center>

## Installation

To use this with your Go project:

```bash
$ go get github.com/Kaamkiya/nanoid-go
```

## Usage

It's a simple library, and it exports only one function and two constants:

```go
const DefaultLength = 22
const DefaultCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890-_"

func Nanoid(length int, charset []string) string
```

To use the defaults:

```go
import "github.com/Kaamkiya/nanoid-go"

nanoid.Nanoid(nanoid.DefaultLength, nanoid.DefaultCharset)
```

More complete examples can be found in [`_examples/`](_examples).

## Features

* Cryptographically secure - it uses a secure random number generator to make
  the IDs.
* Fast.
* Small - the only import is in the standard library.
* URL safe - it only uses URL safe characters, unless you want it to use more.

## License

This project uses the link:LICENSE[Unlicense].
