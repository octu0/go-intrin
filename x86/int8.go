package x86

import _ "unsafe"

//go:linkname Int8Add github.com/octu0/go-intrin/x86.emmAddInt8
//go:noescape
func Int8Add(a, b [16]int8) [16]int8

//go:linkname Int8Add32 github.com/octu0/go-intrin/x86.immAddInt8
//go:noescape
func Int8Add32(a, b [32]int8) [32]int8

//go:linkname Int8Sub github.com/octu0/go-intrin/x86.emmSubInt8
//go:noescape
func Int8Sub(a, b [16]int8) [16]int8

//go:linkname Int8Sub32 github.com/octu0/go-intrin/x86.immSubInt8
//go:noescape
func Int8Sub32(a, b [32]int8) [32]int8

//go:linkname Int8Abs32 github.com/octu0/go-intrin/x86.immAbsInt8
//go:noescape
func Int8Abs32(a [32]int8) [32]int8

//go:linkname Int8Min32 github.com/octu0/go-intrin/x86.immMinInt8
//go:noescape
func Int8Min32(a, b [32]int8) [32]int8

//go:linkname Int8Max32 github.com/octu0/go-intrin/x86.immMaxInt8
//go:noescape
func Int8Max32(a, b [32]int8) [32]int8

func Int8Sum(values ...int8) int8 {
	if 1024 < len(values) {
		return int8SumM256(values)
	}
	return int8SumM128(values)
}

func int8SumM128(values []int8) int8 {
	aligned := alignSlice(values, 16)
	return emmBulkSumInt8(aligned, len(aligned))
}

func int8SumM256(values []int8) int8 {
	aligned := alignSlice(values, 32)
	return immBulkSumInt8(aligned, len(aligned))
}

func int8SumNative(values []int8) int8 {
	sum := int8(0)
	for _, v := range values {
		sum += v
	}
	return sum
}

//go:linkname Uint8Add github.com/octu0/go-intrin/x86.emmAddUint8
//go:noescape
func Uint8Add(a, b [16]uint8) [16]uint8

//go:linkname Uint8Add32 github.com/octu0/go-intrin/x86.immAddUint8
//go:noescape
func Uint8Add32(a, b [32]uint8) [32]uint8

//go:linkname Uint8Sub github.com/octu0/go-intrin/x86.emmSubUint8
//go:noescape
func Uint8Sub(a, b [16]uint8) [16]uint8

//go:linkname Uint8Sub32 github.com/octu0/go-intrin/x86.immSubUint8
//go:noescape
func Uint8Sub32(a, b [32]uint8) [32]uint8

//go:linkname Uint8Avg github.com/octu0/go-intrin/x86.emmAvgUint8
//go:noescape
func Uint8Avg(a, b [16]uint8) [16]uint8

//go:linkname Uint8Avg32 github.com/octu0/go-intrin/x86.immAvgUint8
//go:noescape
func Uint8Avg32(a, b [32]uint8) [32]uint8

//go:linkname Uint8Min github.com/octu0/go-intrin/x86.emmMinUint8
//go:noescape
func Uint8Min(a, b [16]uint8) [16]uint8

//go:linkname Uint8Min32 github.com/octu0/go-intrin/x86.immMinUint8
//go:noescape
func Uint8Min32(a, b [32]uint8) [32]uint8

//go:linkname Uint8Max github.com/octu0/go-intrin/x86.emmMaxUint8
//go:noescape
func Uint8Max(a, b [16]uint8) [16]uint8

//go:linkname Uint8Max32 github.com/octu0/go-intrin/x86.immMaxUint8
//go:noescape
func Uint8Max32(a, b [32]uint8) [32]uint8

func Uint8Sum(values ...uint8) uint8 {
	if 1024 < len(values) {
		return uint8SumM256(values)
	}
	return uint8SumM128(values)
}

func uint8SumM128(values []uint8) uint8 {
	aligned := alignSlice(values, 16)
	return emmBulkSumUint8(aligned, len(aligned))
}

func uint8SumM256(values []uint8) uint8 {
	aligned := alignSlice(values, 32)
	return immBulkSumUint8(aligned, len(aligned))
}

func uint8SumNative(values []uint8) uint8 {
	sum := uint8(0)
	for _, v := range values {
		sum += v
	}
	return sum
}

func Int8ToFloat32(values ...int8) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := xmmBulkConvertInt8ToFloat32(aligned, len(aligned))
	return converted[:initSize]
}

func Uint8ToFloat32(values ...uint8) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := xmmBulkConvertUint8ToFloat32(aligned, len(aligned))
	return converted[:initSize]
}

func Float32ToInt8(values ...float32) []int8 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := xmmBulkConvertFloat32ToInt8(aligned, len(aligned))
	return converted[:initSize]
}

func Float32ToUint8(values ...float32) []uint8 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := xmmBulkConvertFloat32ToUint8(aligned, len(aligned))
	return converted[:initSize]
}
