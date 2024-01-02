package x86

/*
#cgo CFLAGS: -msse2
#include "emm.h"
*/
import "C"

import (
	"unsafe"
)

func EmmAddInt8(a, b [16]int8) [16]int8 {
	out := [16]int8{}
	C.emm_add_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func EmmSubInt8(a, b [16]int8) [16]int8 {
	out := [16]int8{}
	C.emm_sub_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func EmmAddUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_add_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func EmmSubUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_sub_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func EmmAvgUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_avg_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func EmmMaxUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_max_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func EmmMinUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_min_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}
