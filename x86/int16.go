package x86

import _ "unsafe"

//go:linkname Int16Add github.com/octu0/go-intrin/x86.emmAddInt16
//go:noescape
func Int16Add(a, b [8]int16) [8]int16

//go:linkname Int16Sub github.com/octu0/go-intrin/x86.emmSubInt16
//go:noescape
func Int16Sub(a, b [8]int16) [8]int16

//go:linkname Int16Min github.com/octu0/go-intrin/x86.emmMinInt16
//go:noescape
func Int16Min(a, b [8]int16) [8]int16

//go:linkname Int16Max github.com/octu0/go-intrin/x86.emmMaxInt16
//go:noescape
func Int16Max(a, b [8]int16) [8]int16

//go:linkname Int16MAdd github.com/octu0/go-intrin/x86.emmMAddInt16
//go:noescape
func Int16MAdd(a, b [8]int16) [4]int32

func Int16Sum(values ...int16) int16 {
	if len(values) < 256 {
		return int16SumNative(values)
	}

	aligned := alignSlice(values, 8)
	return emmBulkSumInt16(aligned, len(aligned))
}

func int16SumNative(values []int16) int16 {
	sum := int16(0)
	for _, v := range values {
		sum += v
	}
	return sum
}

//go:linkname Uint16Add github.com/octu0/go-intrin/x86.emmAddUint16
//go:noescape
func Uint16Add(a, b [8]uint16) [8]uint16

//go:linkname Uint16Sub github.com/octu0/go-intrin/x86.emmSubUint16
//go:noescape
func Uint16Sub(a, b [8]uint16) [8]uint16

//go:linkname Uint16Avg github.com/octu0/go-intrin/x86.emmAvgUint16
//go:noescape
func Uint16Avg(a, b [8]uint16) [8]uint16

func Uint16Sum(values ...uint16) uint16 {
	if len(values) < 256 {
		return uint16SumNative(values)
	}

	aligned := alignSlice(values, 8)
	return emmBulkSumUint16(aligned, len(aligned))
}

func uint16SumNative(values []uint16) uint16 {
	sum := uint16(0)
	for _, v := range values {
		sum += v
	}
	return sum
}
