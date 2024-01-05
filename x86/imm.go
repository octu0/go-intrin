package x86

/*
#cgo CFLAGS: -mavx
#include "imm.h"
*/
import "C"

import (
	"unsafe"
)

func immAddFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_add_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immSubFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_sub_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immMulFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_mul_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immDivFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_div_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immSqrtFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_sqrt_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
	)
	return out
}

func immMinFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_min_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immMaxFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_max_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAndFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_and_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immOrFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_or_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immXorFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_xor_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAndNotFloat64(a [4]float64, b [4]float64) [4]float64 {
	out := [4]float64{}
	C.imm_andnot_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immBulkConvertInt32ToFloat64(in []int32, size int) []float64 {
	out := make([]float64, size)
	C.imm_bulk_convert_int32_to_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func immBulkConvertFloat64ToInt32(in []float64, size int) []int32 {
	out := make([]int32, size)
	C.imm_bulk_convert_float64_to_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func immAddFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_add_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immSubFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_sub_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immMulFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_mul_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immDivFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_div_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immSqrtFloat32(a [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_sqrt_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
	)
	return out
}

func immRSqrtFloat32(a [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_rsqrt_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
	)
	return out
}

func immRcpFloat32(a [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_rcp_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
	)
	return out
}

func immMinFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_min_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immMaxFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_max_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAndFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_and_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immOrFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_or_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immXorFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_xor_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAndNotFloat32(a [8]float32, b [8]float32) [8]float32 {
	out := [8]float32{}
	C.imm_andnot_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immBulkConvertInt32ToFloat32(in []int32, size int) []float32 {
	out := make([]float32, size)
	C.imm_bulk_convert_int32_to_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func immBulkConvertFloat32ToInt32(in []float32, size int) []int32 {
	out := make([]int32, size)
	C.imm_bulk_convert_float32_to_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}
