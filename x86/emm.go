package x86

/*
#cgo CFLAGS: -msse2
#include "emm.h"
*/
import "C"

import (
	"unsafe"
)

func emmAddInt8(a, b [16]int8) [16]int8 {
	out := [16]int8{}
	C.emm_add_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmSubInt8(a, b [16]int8) [16]int8 {
	out := [16]int8{}
	C.emm_sub_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmAddUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_add_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmSubUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_sub_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmAvgUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_avg_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmMinUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_min_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmMaxUint8(a, b [16]uint8) [16]uint8 {
	out := [16]uint8{}
	C.emm_max_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmAddInt16(a, b [8]int16) [8]int16 {
	out := [8]int16{}
	C.emm_add_int16(
		(*C.int16_t)(unsafe.Pointer(&out[0])),
		(*C.int16_t)(unsafe.Pointer(&a[0])),
		(*C.int16_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmSubInt16(a, b [8]int16) [8]int16 {
	out := [8]int16{}
	C.emm_sub_int16(
		(*C.int16_t)(unsafe.Pointer(&out[0])),
		(*C.int16_t)(unsafe.Pointer(&a[0])),
		(*C.int16_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmMaxInt16(a, b [8]int16) [8]int16 {
	out := [8]int16{}
	C.emm_max_int16(
		(*C.int16_t)(unsafe.Pointer(&out[0])),
		(*C.int16_t)(unsafe.Pointer(&a[0])),
		(*C.int16_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmMinInt16(a, b [8]int16) [8]int16 {
	out := [8]int16{}
	C.emm_min_int16(
		(*C.int16_t)(unsafe.Pointer(&out[0])),
		(*C.int16_t)(unsafe.Pointer(&a[0])),
		(*C.int16_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

//	[4]int32{
//	  (a[0] * b[0]) + (a[1] * b[1]),
//	  (a[2] * b[2]) + (a[3] * b[3]),
//	  (a[4] * b[4]) + (a[5] * b[5]),
//	  (a[6] * b[6]) + (a[7] * b[7]),
//	}
func emmMAddInt16(a, b [8]int16) [4]int32 {
	out := [4]int32{}
	C.emm_madd_int16(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.int16_t)(unsafe.Pointer(&a[0])),
		(*C.int16_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmAddUint16(a, b [8]uint16) [8]uint16 {
	out := [8]uint16{}
	C.emm_add_uint16(
		(*C.uint16_t)(unsafe.Pointer(&out[0])),
		(*C.uint16_t)(unsafe.Pointer(&a[0])),
		(*C.uint16_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmSubUint16(a, b [8]uint16) [8]uint16 {
	out := [8]uint16{}
	C.emm_sub_uint16(
		(*C.uint16_t)(unsafe.Pointer(&out[0])),
		(*C.uint16_t)(unsafe.Pointer(&a[0])),
		(*C.uint16_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmAvgUint16(a, b [8]uint16) [8]uint16 {
	out := [8]uint16{}
	C.emm_avg_uint16(
		(*C.uint16_t)(unsafe.Pointer(&out[0])),
		(*C.uint16_t)(unsafe.Pointer(&a[0])),
		(*C.uint16_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmAddFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_add_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmSubFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_sub_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmMulFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_mul_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmDivFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_div_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmSqrtFloat64(a [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_sqrt_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
	)
	return out
}

func emmMaxFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_max_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmMinFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_min_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmAndFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_and_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmOrFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_or_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmXorFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_or_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmAndNotFloat64(a, b [2]float64) [2]float64 {
	out := [2]float64{}
	C.emm_andnot_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&a[0])),
		(*C.double)(unsafe.Pointer(&b[0])),
	)
	return out
}

func emmBulkSumInt8(in []int8, size int) int8 {
	out := [16]int8{}
	C.emm_bulk_sum_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out[0] + out[1] + out[2] + out[3] +
		out[4] + out[5] + out[6] + out[7] +
		out[8] + out[9] + out[10] + out[11] +
		out[12] + out[13] + out[14] + out[15]
}

func emmBulkSumUint8(in []uint8, size int) uint8 {
	out := [16]uint8{}
	C.emm_bulk_sum_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out[0] + out[1] + out[2] + out[3] +
		out[4] + out[5] + out[6] + out[7] +
		out[8] + out[9] + out[10] + out[11] +
		out[12] + out[13] + out[14] + out[15]
}

func emmBulkSumInt16(in []int16, size int) int16 {
	out := [8]int16{}
	C.emm_bulk_sum_int16(
		(*C.int16_t)(unsafe.Pointer(&out[0])),
		(*C.int16_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out[0] + out[1] + out[2] + out[3] +
		out[4] + out[5] + out[6] + out[7]
}

func emmBulkSumUint16(in []uint16, size int) uint16 {
	out := [8]uint16{}
	C.emm_bulk_sum_uint16(
		(*C.uint16_t)(unsafe.Pointer(&out[0])),
		(*C.uint16_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out[0] + out[1] + out[2] + out[3] +
		out[4] + out[5] + out[6] + out[7]
}

func emmBulkConvertInt32ToFloat32(in []int32, size int) []float32 {
	out := make([]float32, size)
	C.emm_bulk_convert_int32_to_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func emmBulkConvertFloat32ToInt32(in []float32, size int) []int32 {
	out := make([]int32, size)
	C.emm_bulk_convert_float32_to_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func emmBulkConvertInt32ToFloat64(in []int32, size int) []float64 {
	out := make([]float64, size)
	C.emm_bulk_convert_int32_to_float64(
		(*C.double)(unsafe.Pointer(&out[0])),
		(*C.int32_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func emmBulkConvertFloat64ToInt32(in []float64, size int) []int32 {
	out := make([]int32, size)
	C.emm_bulk_convert_float64_to_int32(
		(*C.int32_t)(unsafe.Pointer(&out[0])),
		(*C.double)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}
