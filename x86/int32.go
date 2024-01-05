package x86

import _ "unsafe"

//go:linkname Int32Add github.com/octu0/go-intrin/x86.immAddInt32
//go:noescape
func Int32Add(a, b [8]int32) [8]int32

//go:linkname Int32Sub github.com/octu0/go-intrin/x86.immSubInt32
//go:noescape
func Int32Sub(a, b [8]int32) [8]int32

//go:linkname Int32Abs github.com/octu0/go-intrin/x86.immAbsInt32
//go:noescape
func Int32Abs(a [8]int32) [8]int32

//go:linkname Int32And github.com/octu0/go-intrin/x86.immAndInt32
//go:noescape
func Int32And(a, b [8]int32) [8]int32

//go:linkname Int32Or github.com/octu0/go-intrin/x86.immOrInt32
//go:noescape
func Int32Or(a, b [8]int32) [8]int32

//go:linkname Int32Xor github.com/octu0/go-intrin/x86.immXorInt32
//go:noescape
func Int32Xor(a, b [8]int32) [8]int32

//go:linkname Int32AndNot github.com/octu0/go-intrin/x86.immAndNotInt32
//go:noescape
func Int32AndNot(a, b [8]int32) [8]int32

func Int32ToFloat64(values ...int32) []float64 {
	if 1024 < len(values) {
		return int32ToFloat64M256(values)
	}
	return int32ToFloat64M128(values)
}

func int32ToFloat64M128(values []int32) []float64 {
	initSize := len(values)
	aligned := alignSlice(values, 2)
	converted := emmBulkConvertInt32ToFloat64(aligned, len(aligned))
	return converted[:initSize]
}

func int32ToFloat64M256(values []int32) []float64 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := immBulkConvertInt32ToFloat64(aligned, len(aligned))
	return converted[:initSize]
}

func Float64ToInt32(values ...float64) []int32 {
	if 1024 < len(values) {
		return float64ToInt32M256(values)
	}
	return float64ToInt32M128(values)
}

func float64ToInt32M128(values []float64) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 2)
	converted := emmBulkConvertFloat64ToInt32(aligned, len(aligned))
	return converted[:initSize]
}

func float64ToInt32M256(values []float64) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := immBulkConvertFloat64ToInt32(aligned, len(aligned))
	return converted[:initSize]
}

func Int32ToFloat32(values ...int32) []float32 {
	if 1024 < len(values) {
		return int32ToFloat32M256(values)
	}
	return int32ToFloat32M128(values)
}

func int32ToFloat32M128(values []int32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := emmBulkConvertInt32ToFloat32(aligned, len(aligned))
	return converted[:initSize]
}

func int32ToFloat32M256(values []int32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 8)
	converted := immBulkConvertInt32ToFloat32(aligned, len(aligned))
	return converted[:initSize]
}

func Float32ToInt32(values ...float32) []int32 {
	if 1024 < len(values) {
		return float32ToInt32M256(values)
	}
	return float32ToInt32M128(values)
}

func float32ToInt32M128(values []float32) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := emmBulkConvertFloat32ToInt32(aligned, len(aligned))
	return converted[:initSize]
}

func float32ToInt32M256(values []float32) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 8)
	converted := immBulkConvertFloat32ToInt32(aligned, len(aligned))
	return converted[:initSize]
}
