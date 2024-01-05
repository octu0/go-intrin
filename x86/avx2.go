package x86

/*
#cgo CFLAGS: -mavx2
#include "avx2.h"
*/
import "C"

import (
	"unsafe"
)

func immAddInt32(a [8]int32, b [8]int32) [8]int32 {
	out := [8]int32{}
	C.imm_add_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&a[0])),
		(*C.int32_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immSubInt32(a [8]int32, b [8]int32) [8]int32 {
	out := [8]int32{}
	C.imm_sub_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&a[0])),
		(*C.int32_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAbsInt32(a [8]int32) [8]int32 {
	out := [8]int32{}
	C.imm_abs_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&a[0])),
	)
	return out
}

func immAndInt32(a [8]int32, b [8]int32) [8]int32 {
	out := [8]int32{}
	C.imm_and_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&a[0])),
		(*C.int32_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immOrInt32(a [8]int32, b [8]int32) [8]int32 {
	out := [8]int32{}
	C.imm_or_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&a[0])),
		(*C.int32_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immXorInt32(a [8]int32, b [8]int32) [8]int32 {
	out := [8]int32{}
	C.imm_xor_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&a[0])),
		(*C.int32_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAndNotInt32(a [8]int32, b [8]int32) [8]int32 {
	out := [8]int32{}
	C.imm_andnot_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&a[0])),
		(*C.int32_t)(unsafe.Pointer(&b[0])),
	)
	return out
}
