//go:build fma

package x86

/*
#cgo CFLAGS: -mavx2 -mfma
#include "avx2_fma.h"
*/
import "C"

import (
	"unsafe"
)

func immFMAddFloat32_128(a, b, c [4]float32) [4]float32 {
	out := [4]float32{}
	C.imm_fmadd_float32_128(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(*C.float)(unsafe.Pointer(&c[0])),
	)
	return out
}

func immFMAddFloat32_256(a, b, c [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_fmadd_float32_256(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(*C.float)(unsafe.Pointer(&c[0])),
	)
	return out
}

func immFMSubFloat32_128(a, b, c [4]float32) [4]float32 {
	out := [4]float32{}
	C.imm_fmsub_float32_128(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(*C.float)(unsafe.Pointer(&c[0])),
	)
	return out
}

func immFMSubFloat32_256(a, b, c [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_fmsub_float32_256(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(*C.float)(unsafe.Pointer(&c[0])),
	)
	return out
}

func immFMAddFloat64_128(a, b, c [2]float64) [2]float64 {
	out := [2]float64{}
	C.imm_fmadd_float64_128(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
		(*C.double)(unsafe.Pointer(&c[0])),
	)
	return out
}

func immFMAddFloat64_256(a, b, c [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_fmadd_float64_256(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
		(*C.double)(unsafe.Pointer(&c[0])),
	)
	return out
}

func immFMSubFloat64_128(a, b, c [2]float64) [2]float64 {
	out := [2]float64{}
	C.imm_fmsub_float64_128(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
		(*C.double)(unsafe.Pointer(&c[0])),
	)
	return out
}

func immFMSubFloat64_256(a, b, c [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_fmsub_float64_256(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
		(*C.double)(unsafe.Pointer(&c[0])),
	)
	return out
}
