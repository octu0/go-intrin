package x86

/*
#cgo CFLAGS: -msse
#include "xmm.h"
*/
import "C"

import (
	"unsafe"
)

func XmmAdd(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_add(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmSub(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_sub(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmMul(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_mul(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmDiv(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_div(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmSqrt(a [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_sqrt(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
	)
	return out
}

func XmmRSqrt(a [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_rsqrt(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
	)
	return out
}

func XmmRcp(a [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_rcp(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
	)
	return out
}

func XmmMin(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_min(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmMax(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_max(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmAnd(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_and(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmOr(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_or(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmXor(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_xor(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmAndNot(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_andnot(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func XmmBulkConvertFromInt8(in []int8, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_convert_from_int8(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func XmmBulkConvertFromUint8(in []uint8, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_convert_from_uint8(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func XmmBulkConvertToInt8(in []float32, size int) []int8 {
	out := make([]int8, size)
	C.xmm_bulk_convert_to_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func XmmBulkConvertToUint8(in []float32, size int) []uint8 {
	out := make([]uint8, size)
	C.xmm_bulk_convert_to_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func XmmBulkSum(a []float32) float32 {
	out := [4]float32{}
	C.xmm_bulk_sum(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		C.int(len(a)),
	)
	return out[0] + out[1] + out[2] + out[3]
}

func XmmBulkSum2(out, a, b []float32, size int) {
	C.xmm_bulk_sum2(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		C.int(size),
	)
}

func XmmBulkAdd(base [4]float32, in []float32, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_add(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func XmmBulkSub(base [4]float32, in []float32, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_sub(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func XmmBulkMul(base [4]float32, in []float32, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_mul(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func XmmBulkDiv(base [4]float32, in []float32, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_div(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func XmmTile4x4Sum(in []float32, size int) []float32 {
	out := make([]float32, size/4)
	C.xmm_tile4x4_sum(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}
