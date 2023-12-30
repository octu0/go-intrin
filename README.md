# `go-intrin`

[![MIT License](https://img.shields.io/github/license/octu0/go-intrin)](https://github.com/octu0/go-intrin/blob/master/LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/octu0/go-intrin)](https://pkg.go.dev/github.com/octu0/go-intrin)
[![Go Report Card](https://goreportcard.com/badge/github.com/octu0/go-intrin)](https://goreportcard.com/report/github.com/octu0/go-intrin)
[![Releases](https://img.shields.io/github/v/release/octu0/go-intrin)](https://github.com/octu0/go-intrin/releases)

`go-intrin` is a Go library that provides a simple interface to SIMD (MMX/SSE/SSE2/AVX/AVX2) instructions.
currently supports Intel x86.

## Installation

```bash
go get github.com/octu0/go-intrin
```

## Example

### float32

```go
import "github.com/octu0/go-intrin/x86"

func Add() {
    a := [4]float32{ 1.0, 3.0, 5.0, 11.0 }
    b := [4]float32{ 2.0, 4.0, 6.0, 8.0 }

    r := x86.Float32Add(a, b)
    r[0] // 3.0
    r[1] // 7.0
    r[2] // 11.0
    r[3] // 19.0
}

func Sub() {
    a := [4]float32{ 1.0, 3.0, -5.0, -11.0 }
    b := [4]float32{ 2.0, -4.0, 6.0, -8.0 }

    r := x86.Float32Sub(a, b)
    r[0] // 3.0
    r[1] // 7.0
    r[2] // -11.0
    r[3] // -3
}
```

# License

MIT, see LICENSE file for details.
