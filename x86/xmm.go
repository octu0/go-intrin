package x86

/*
#cgo CFLAGS: -msse
#include "xmm.h"
*/
import "C"

import (
	"unsafe"
)

//go:linkname __mm_setr_ps xmm_setr_ps
//go:noescape
func __mm_setr_ps(C.float, C.float, C.float, C.float) C.__m128

//go:linkname __mm_add_ps xmm_add_ps
//go:noescape
func __mm_add_ps(C.__m128, C.__m128) C.__m128

//go:linkname __mm_store_ps xmm_store_ps
//go:noescape
func __mm_store_ps(*C.float, C.__m128)

func __xmmAdd(a, b [4]float32) [4]float32 {
	ma := __mm_setr_ps(
		C.float(a[0]),
		C.float(a[1]),
		C.float(a[2]),
		C.float(a[3]),
	)
	mb := __mm_setr_ps(
		C.float(b[0]),
		C.float(b[1]),
		C.float(b[2]),
		C.float(b[3]),
	)
	r := __mm_add_ps(ma, mb)

	out := [4]float32{}
	__mm_store_ps(
		(*C.float)(unsafe.Pointer(&out[0])),
		r,
	)
	return out
}

func xmmAdd(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_add(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmSub(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_sub(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmMul(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_mul(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmDiv(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_div(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmSqrt(a [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_sqrt(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
	)
	return out
}

func xmmRSqrt(a [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_rsqrt(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
	)
	return out
}

func xmmRcp(a [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_rcp(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
	)
	return out
}

func xmmMin(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_min(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmMax(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_max(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmAnd(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_and(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmOr(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_or(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmXor(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_xor(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmAndNot(a, b [4]float32) [4]float32 {
	out := [4]float32{}
	C.xmm_andnot(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func xmmBulkConvertInt8ToFloat32(in []int8, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_convert_int8_to_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func xmmBulkConvertUint8ToFloat32(in []uint8, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_convert_uint8_to_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func xmmBulkConvertFloat32ToInt8(in []float32, size int) []int8 {
	out := make([]int8, size)
	C.xmm_bulk_convert_float32_to_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func xmmBulkConvertFloat32ToUint8(in []float32, size int) []uint8 {
	out := make([]uint8, size)
	C.xmm_bulk_convert_float32_to_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func xmmBulkSum(a []float32, size int) float32 {
	out := [4]float32{}
	C.xmm_bulk_sum(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		C.int(size),
	)
	return out[0] + out[1] + out[2] + out[3]
}

func xmmBulkSum2(out, a, b []float32, size int) {
	C.xmm_bulk_sum2(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		C.int(size),
	)
}

func xmmBulkAdd(base [4]float32, in []float32, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_add(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func xmmBulkSub(base [4]float32, in []float32, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_sub(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func xmmBulkMul(base [4]float32, in []float32, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_mul(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func xmmBulkDiv(base [4]float32, in []float32, size int) []float32 {
	out := make([]float32, size)
	C.xmm_bulk_div(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&base[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func xmmTile4x4Sum(in []float32, size int) []float32 {
	out := make([]float32, size/4)
	C.xmm_tile4x4_sum(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func xmmRGBAGrayscale(in []byte, size int) []byte {
	out := make([]byte, size)
	C.xmm_grayscale(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}
