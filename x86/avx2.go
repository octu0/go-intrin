package x86

/*
#cgo CFLAGS: -mavx2
#include "avx2.h"
*/
import "C"

import (
	"unsafe"
)

func immAddInt8(a [32]int8, b [32]int8) [32]int8 {
	out := [32]int8{}
	C.imm_add_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immSubInt8(a [32]int8, b [32]int8) [32]int8 {
	out := [32]int8{}
	C.imm_sub_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAbsInt8(a [32]int8) [32]int8 {
	out := [32]int8{}
	C.imm_abs_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
	)
	return out
}

func immMinInt8(a [32]int8, b [32]int8) [32]int8 {
	out := [32]int8{}
	C.imm_min_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immMaxInt8(a [32]int8, b [32]int8) [32]int8 {
	out := [32]int8{}
	C.imm_max_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAndInt8(a [32]int8, b [32]int8) [32]int8 {
	out := [32]int8{}
	C.imm_and_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immOrInt8(a [32]int8, b [32]int8) [32]int8 {
	out := [32]int8{}
	C.imm_or_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immXorInt8(a [32]int8, b [32]int8) [32]int8 {
	out := [32]int8{}
	C.imm_xor_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAndNotInt8(a [32]int8, b [32]int8) [32]int8 {
	out := [32]int8{}
	C.imm_andnot_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&a[0])),
		(*C.int8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAddUint8(a [32]uint8, b [32]uint8) [32]uint8 {
	out := [32]uint8{}
	C.imm_add_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immSubUint8(a [32]uint8, b [32]uint8) [32]uint8 {
	out := [32]uint8{}
	C.imm_sub_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAvgUint8(a [32]uint8, b [32]uint8) [32]uint8 {
	out := [32]uint8{}
	C.imm_avg_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immMinUint8(a [32]uint8, b [32]uint8) [32]uint8 {
	out := [32]uint8{}
	C.imm_min_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immMaxUint8(a [32]uint8, b [32]uint8) [32]uint8 {
	out := [32]uint8{}
	C.imm_max_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAndUint8(a [32]uint8, b [32]uint8) [32]uint8 {
	out := [32]uint8{}
	C.imm_and_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immOrUint8(a [32]uint8, b [32]uint8) [32]uint8 {
	out := [32]uint8{}
	C.imm_or_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immXorUint8(a [32]uint8, b [32]uint8) [32]uint8 {
	out := [32]uint8{}
	C.imm_xor_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

func immAndNotUint8(a [32]uint8, b [32]uint8) [32]uint8 {
	out := [32]uint8{}
	C.imm_andnot_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&a[0])),
		(*C.uint8_t)(unsafe.Pointer(&b[0])),
	)
	return out
}

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

func immBulkSumInt8(in []int8, size int) int8 {
	out := [32]int8{}
	C.imm_bulk_sum_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out[0] + out[1] + out[2] + out[3] +
		out[4] + out[5] + out[6] + out[7] +
		out[8] + out[9] + out[10] + out[11] +
		out[12] + out[13] + out[14] + out[15] +
		out[16] + out[17] + out[18] + out[19] +
		out[20] + out[21] + out[22] + out[23] +
		out[24] + out[25] + out[26] + out[27] +
		out[28] + out[29] + out[30] + out[31]
}

func immBulkSumUint8(in []uint8, size int) uint8 {
	out := [32]uint8{}
	C.imm_bulk_sum_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out[0] + out[1] + out[2] + out[3] +
		out[4] + out[5] + out[6] + out[7] +
		out[8] + out[9] + out[10] + out[11] +
		out[12] + out[13] + out[14] + out[15] +
		out[16] + out[17] + out[18] + out[19] +
		out[20] + out[21] + out[22] + out[23] +
		out[24] + out[25] + out[26] + out[27] +
		out[28] + out[29] + out[30] + out[31]
}

func immBulkConvertInt8ToFloat32(in []int8, size int) []float32 {
	out := make([]float32, size)
	C.imm_bulk_convert_int8_to_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.int8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func immBulkConvertUint8ToFloat32(in []uint8, size int) []float32 {
	out := make([]float32, size)
	C.imm_bulk_convert_uint8_to_float32(
		(*C.float)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func immBulkConvertFloat32ToInt8(in []float32, size int) []int8 {
	out := make([]int8, size)
	C.imm_bulk_convert_float32_to_int8(
		(*C.int8_t)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func immBulkConvertFloat32ToUint8(in []float32, size int) []uint8 {
	out := make([]uint8, size)
	C.imm_bulk_convert_float32_to_uint8(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.float)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}

func immRGBAGrayscale(in []byte, size int) []byte {
	out := make([]byte, size)
	C.imm_grayscale(
		(*C.uint8_t)(unsafe.Pointer(&out[0])),
		(*C.uint8_t)(unsafe.Pointer(&in[0])),
		C.int(size),
	)
	return out
}
