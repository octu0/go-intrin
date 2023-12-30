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
