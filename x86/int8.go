package x86

import _ "unsafe"

//go:linkname Int8Add github.com/octu0/go-intrin/x86.emmAddInt8
//go:noescape
func Int8Add(a, b [16]int8) [16]int8

//go:linkname Int8Sub github.com/octu0/go-intrin/x86.emmSubInt8
//go:noescape
func Int8Sub(a, b [16]int8) [16]int8

func Int8Sum(values ...int8) int8 {
	if len(values) < 128 {
		return int8SumNative(values)
	}

	aligned := alignSlice(values, 16)
	return emmBulkSumInt8(aligned, len(aligned))
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

//go:linkname Uint8Sub github.com/octu0/go-intrin/x86.emmSubUint8
//go:noescape
func Uint8Sub(a, b [16]uint8) [16]uint8

//go:linkname Uint8Avg github.com/octu0/go-intrin/x86.emmAvgUint8
//go:noescape
func Uint8Avg(a, b [16]uint8) [16]uint8

//go:linkname Uint8Min github.com/octu0/go-intrin/x86.emmMinUint8
//go:noescape
func Uint8Min(a, b [16]uint8) [16]uint8

//go:linkname Uint8Max github.com/octu0/go-intrin/x86.emmMaxUint8
//go:noescape
func Uint8Max(a, b [16]uint8) [16]uint8

func Uint8Sum(values ...uint8) uint8 {
	if len(values) < 128 {
		return uint8SumNative(values)
	}

	aligned := alignSlice(values, 16)
	return emmBulkSumUint8(aligned, len(aligned))
}

func uint8SumNative(values []uint8) uint8 {
	sum := uint8(0)
	for _, v := range values {
		sum += v
	}
	return sum
}
