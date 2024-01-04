# `go-intrin`

[![MIT License](https://img.shields.io/github/license/octu0/go-intrin)](https://github.com/octu0/go-intrin/blob/master/LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/octu0/go-intrin)](https://pkg.go.dev/github.com/octu0/go-intrin)
[![Go Report Card](https://goreportcard.com/badge/github.com/octu0/go-intrin)](https://goreportcard.com/report/github.com/octu0/go-intrin)
[![Releases](https://img.shields.io/github/v/release/octu0/go-intrin)](https://github.com/octu0/go-intrin/releases)

`go-intrin` is a Go library that provides a simple interface to SIMD (MMX/SSE/SSE2) instructions.
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

### Int8/Int16/Uint8/Uint16

```go
import "github.com/octu0/go-intrin/x86"

func Add() {
    a := [16]int8{
      1, 2, 3, 4, 5, 6, 7, 8,
      -1, 2, 3, -4, -5, -6, -7, -8,
    }
    b := [16]int8{
      11, 12, 13, 14, 15, 16, 17, 18,
      100, 125, 126, -128, -124, 127, -2, -3,
    }
    r := x86.Int8Add(a, b)
}

func Sub() {
    a := [8]int16{1, 2, 32767, -32768, 1, 1, -32768, 0}
    b := [8]int16{1, 3, 1, 1, -32768, 32767, -1, -1}
    r := x86.Int16Add(a, b)
}

func Max() {
    a := [16]uint8{
      1, 2, 3, 4, 5, 6, 7, 8,
      255, 250, 0, 255, 127, 6, 7, 8,
    }
    b := [16]uint8{
      0, 1, 10, 255, 254, 0, 7, 8,
      255, 251, 255, 1, 128, 7, 6, 9,
    }
    r := x86.Uint8Max(a, b)
}

func Avg() {
    a := [8]uint16{1, 2, 32767, 65535, 1, 65535, 0, 0}
    b := [8]uint16{1, 3, 1, 1, 65535, 65535, 0, 1}
    r := x86.Uint8Avg(a, b)
}
```

## Benchmark

The performance of SIMD benchmarks can vary depending on the execution environment.  
In some cases, performance can be better than native Go even when using cgo to access SIMD instructions.

Darwin & Core i5

```
goos: darwin
goarch: amd64
pkg: github.com/octu0/go-intrin/x86
cpu: Intel(R) Core(TM) i5-8210Y CPU @ 1.60GHz
BenchmarkInt16
BenchmarkInt16/sum/go
BenchmarkInt16/sum/go-4         	1000000000	         0.3100 ns/op
BenchmarkInt16/sum/simd
BenchmarkInt16/sum/simd-4       	1000000000	         0.1590 ns/op
BenchmarkInt8
BenchmarkInt8/sum/go
BenchmarkInt8/sum/go-4          	1000000000	         0.3011 ns/op
BenchmarkInt8/sum/simd
BenchmarkInt8/sum/simd-4        	1000000000	         0.08489 ns/op
```

Linux & Xeon

```
$ go test -bench=BenchmarkFloat32 -v ./x86/
goos: linux
goarch: amd64
pkg: github.com/octu0/go-intrin/x86
cpu: Intel(R) Xeon(R) W-11955M CPU @ 2.60GHz
BenchmarkFloat32
BenchmarkFloat32/sum/go
BenchmarkFloat32/sum/go-16         	585412603	         2.097 ns/op
BenchmarkFloat32/sum/simd
BenchmarkFloat32/sum/simd-16       	1000000000	         0.8212 ns/op
```

Go code for Grayscale can improve performance by minimizing cgo calls.

```
goos: linux
goarch: amd64
pkg: github.com/octu0/go-intrin/x86
cpu: Intel(R) Xeon(R) W-11955M CPU @ 2.60GHz
BenchmarkGrayscale
BenchmarkGrayscale/go
BenchmarkGrayscale/go-16           1923    596976 ns/op
BenchmarkGrayscale/simd/small
BenchmarkGrayscale/simd/small-16    100  10063526 ns/op
BenchmarkGrayscale/simd/medium
BenchmarkGrayscale/simd/medium-16  1087   1112508 ns/op
BenchmarkGrayscale/simd/full
BenchmarkGrayscale/simd/full-16    3070    382772 ns/op
PASS
```

The Go code is as below:

```go
func Grayscale(src *image.NRGBA) *image.NRGBA {
  b := src.Bounds()
  w, h := b.Dx(), b.Dy()

  out := image.NewNRGBA(b)
  for y := 0; y < h; y += 1 {
    for x := 0; x < w; x += 1 {
      c := src.NRGBAAt(x, y)
      // BT.709
      gray := byte((0.2126 * float32(c.R)) + (0.7152 * float32(c.G)) + (0.0722 * float32(c.B)))
      out.SetNRGBA(x, y, color.NRGBA{
        R: gray,
        G: gray,
        B: gray,
        A: 0xff,
      })
    }
  }
  return out
}
```

It is called as `C.simd_grayscale` as follows (this code is used in the simd/full benchmark):

```c
static void simd_grayscale(uint8_t *out, uint8_t *in, int size) {
  __m128 bt709 = _mm_setr_ps(0.2126f, 0.7152f, 0.0722f, 0.0f);
  uint8_t gray[8];
  for(int i = 0; i < size; i += 16) {
    __m64 m1 = _mm_setr_pi8(
      in[i + 0], in[i + 1], in[i + 2], in[i + 3],
      0, 0, 0, 0
    );
    __m64 m2 = _mm_setr_pi8(
      in[i + 4], in[i + 5], in[i + 6], in[i + 7],
      0, 0, 0, 0
    );
    __m64 m3 = _mm_setr_pi8(
      in[i + 8], in[i + 9], in[i + 10], in[i + 11],
      0, 0, 0, 0
    );
    __m64 m4 = _mm_setr_pi8(
      in[i + 12], in[i + 13], in[i + 14], in[i + 15],
      0, 0, 0, 0
    );
    __m128 rgba1 = _mm_mul_ps(_mm_cvtpu8_ps(m1), bt701);
    __m128 rgba2 = _mm_mul_ps(_mm_cvtpu8_ps(m2), bt701);
    __m128 rgba3 = _mm_mul_ps(_mm_cvtpu8_ps(m3), bt701);
    __m128 rgba4 = _mm_mul_ps(_mm_cvtpu8_ps(m4), bt701);

    __m128 r = _mm_setr_ps(rgba1[0], rgba2[0], rgba3[0], rgba4[0]); // R
    __m128 g = _mm_setr_ps(rgba1[1], rgba2[1], rgba3[1], rgba4[1]); // G
    __m128 b = _mm_setr_ps(rgba1[2], rgba2[2], rgba3[2], rgba4[2]); // B

    // gray = [rgba1,rgba2,rgba3,rgba4]
    __m128 gray_float = _mm_add_ps(_mm_add_ps(r, g), b);

    __m64 gray_u8 = _mm_cvtps_pi8(gray_float);
    memcpy(&gray, &gray_u8, sizeof(__m64));

    uint8_t tmp[16] = {
      gray[0], gray[0], gray[0], 255,
      gray[1], gray[1], gray[1], 255,
      gray[2], gray[2], gray[2], 255,
      gray[3], gray[3], gray[3], 255
    };
    memcpy(out + i, &tmp, 16);
  }
}
```

# License

MIT, see LICENSE file for details.
