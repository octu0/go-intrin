package x86

import _ "unsafe"

//go:linkname Float64Add github.com/octu0/go-intrin/x86.emmAddFloat64
//go:noescape
func Float64Add(a, b [2]float64) [2]float64

//go:linkname Float64Sub github.com/octu0/go-intrin/x86.emmSubFloat64
//go:noescape
func Float64Sub(a, b [2]float64) [2]float64

//go:linkname Float64Mul github.com/octu0/go-intrin/x86.emmMulFloat64
//go:noescape
func Float64Mul(a, b [2]float64) [2]float64

//go:linkname Float64Div github.com/octu0/go-intrin/x86.emmDivFloat64
//go:noescape
func Float64Div(a, b [2]float64) [2]float64

//go:linkname Float64Sqrt github.com/octu0/go-intrin/x86.emmSqrtFloat64
//go:noescape
func Float64Sqrt(a [2]float64) [2]float64

//go:linkname Float64Min github.com/octu0/go-intrin/x86.emmMinFloat64
//go:noescape
func Float64Min(a, b [2]float64) [2]float64

//go:linkname Float64Max github.com/octu0/go-intrin/x86.emmMaxFloat64
//go:noescape
func Float64Max(a, b [2]float64) [2]float64

//go:linkname Float64And github.com/octu0/go-intrin/x86.emmAndFloat64
//go:noescape
func Float64And(a, b [2]float64) [2]float64

//go:linkname Float64Or github.com/octu0/go-intrin/x86.emmOrFloat64
//go:noescape
func Float64Or(a, b [2]float64) [2]float64

//go:linkname Float64Xor github.com/octu0/go-intrin/x86.emmXorFloat64
//go:noescape
func Float64Xor(a, b [2]float64) [2]float64

//go:linkname Float64AndNot github.com/octu0/go-intrin/x86.emmAndNotFloat64
//go:noescape
func Float64AndNot(a, b [2]float64) [2]float64

func Int32ToFloat64(values ...int32) []float64 {
	initSize := len(values)
	aligned := alignSlice(values, 2)
	converted := emmBulkConvertInt32ToFloat64(aligned, len(aligned))
	return converted[:initSize]
}

func Float64ToInt32(values ...float64) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 2)
	converted := emmBulkConvertFloat64ToInt32(aligned, len(aligned))
	return converted[:initSize]
}
