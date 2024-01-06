package x86

import _ "unsafe"

//go:linkname Float32Add github.com/octu0/go-intrin/x86.xmmAdd
//go:noescape
func Float32Add(a, b [4]float32) [4]float32

//go:linkname Float32Add8 github.com/octu0/go-intrin/x86.immAddFloat32
//go:noescape
func Float32Add8(a, b [8]float32) [8]float32

//go:linkname Float32Sub github.com/octu0/go-intrin/x86.xmmSub
//go:noescape
func Float32Sub(a, b [4]float32) [4]float32

//go:linkname Float32Sub8 github.com/octu0/go-intrin/x86.immSubFloat32
//go:noescape
func Float32Sub8(a, b [8]float32) [8]float32

//go:linkname Float32Mul github.com/octu0/go-intrin/x86.xmmMul
//go:noescape
func Float32Mul(a, b [4]float32) [4]float32

//go:linkname Float32Mul8 github.com/octu0/go-intrin/x86.immMulFloat32
//go:noescape
func Float32Mul8(a, b [8]float32) [8]float32

//go:linkname Float32Div github.com/octu0/go-intrin/x86.xmmDiv
//go:noescape
func Float32Div(a, b [4]float32) [4]float32

//go:linkname Float32Div8 github.com/octu0/go-intrin/x86.immDivFloat32
//go:noescape
func Float32Div8(a, b [8]float32) [8]float32

//go:linkname Float32Sqrt github.com/octu0/go-intrin/x86.xmmSqrt
//go:noescape
func Float32Sqrt(a [4]float32) [4]float32

//go:linkname Float32Sqrt8 github.com/octu0/go-intrin/x86.immSqrtFloat32
//go:noescape
func Float32Sqrt8(a [8]float32) [8]float32

//go:linkname Float32RSqrt github.com/octu0/go-intrin/x86.xmmRSqrt
//go:noescape
func Float32RSqrt(a [4]float32) [4]float32

//go:linkname Float32RSqrt8 github.com/octu0/go-intrin/x86.immRSqrtFloat32
//go:noescape
func Float32RSqrt8(a [8]float32) [8]float32

//go:linkname Float32Rcp github.com/octu0/go-intrin/x86.xmmRcp
//go:noescape
func Float32Rcp(a [4]float32) [4]float32

//go:linkname Float32Rcp8 github.com/octu0/go-intrin/x86.immRcpFloat32
//go:noescape
func Float32Rcp8(a [8]float32) [8]float32

//go:linkname Float32Min github.com/octu0/go-intrin/x86.xmmMin
//go:noescape
func Float32Min(a, b [4]float32) [4]float32

//go:linkname Float32Min8 github.com/octu0/go-intrin/x86.immMinFloat32
//go:noescape
func Float32Min8(a, b [8]float32) [8]float32

//go:linkname Float32Max github.com/octu0/go-intrin/x86.xmmMax
//go:noescape
func Float32Max(a, b [4]float32) [4]float32

//go:linkname Float32Max8 github.com/octu0/go-intrin/x86.immMaxFloat32
//go:noescape
func Float32Max8(a, b [8]float32) [8]float32

//go:linkname Float32And github.com/octu0/go-intrin/x86.xmmAnd
//go:noescape
func Float32And(a, b [4]float32) [4]float32

//go:linkname Float32And8 github.com/octu0/go-intrin/x86.immAndFloat32
//go:noescape
func Float32And8(a, b [8]float32) [8]float32

//go:linkname Float32Or github.com/octu0/go-intrin/x86.xmmOr
//go:noescape
func Float32Or(a, b [4]float32) [4]float32

//go:linkname Float32Or8 github.com/octu0/go-intrin/x86.immOrFloat32
//go:noescape
func Float32Or8(a, b [8]float32) [8]float32

//go:linkname Float32Xor github.com/octu0/go-intrin/x86.xmmXor
//go:noescape
func Float32Xor(a, b [4]float32) [4]float32

//go:linkname Float32Xor8 github.com/octu0/go-intrin/x86.immXorFloat32
//go:noescape
func Float32Xor8(a, b [8]float32) [8]float32

//go:linkname Float32AndNot github.com/octu0/go-intrin/x86.xmmAndNot
//go:noescape
func Float32AndNot(a, b [4]float32) [4]float32

//go:linkname Float32AndNot8 github.com/octu0/go-intrin/x86.immAndNotFloat32
//go:noescape
func Float32AndNot8(a, b [8]float32) [8]float32

func Float32Tile4Sum(values ...float32) []float32 {
	aligned := alignSlice(values, 16)
	tiled := xmmTile4x4Sum(aligned, len(aligned))
	return tiled
}

func Float32Sum(values ...float32) float32 {
	if len(values) < 256 {
		return float32SumNative(values)
	}

	aligned := alignSlice(values, 4)
	return xmmBulkSum(aligned, len(aligned))
}

func float32SumNative(values []float32) float32 {
	sum := float32(0.0)
	for _, v := range values {
		sum += v
	}
	return sum
}

type Float32Filter struct {
	Base [4]float32
}

func (f *Float32Filter) Add(values ...float32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	res := xmmBulkAdd(f.Base, aligned, len(aligned))
	return res[:initSize]
}

func (f *Float32Filter) Sub(values ...float32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	res := xmmBulkSub(f.Base, aligned, len(aligned))
	return res[:initSize]
}

func (f *Float32Filter) Mul(values ...float32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	res := xmmBulkMul(f.Base, aligned, len(aligned))
	return res[:initSize]
}

func (f *Float32Filter) Div(values ...float32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	res := xmmBulkDiv(f.Base, aligned, len(aligned))
	return res[:initSize]
}
