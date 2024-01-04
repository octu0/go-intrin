package x86

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestEmmAddInt8(t *testing.T) {
	out := EmmAddInt8(
		[16]int8{
			1, 2, 3, 4, 5, 6, 7, 8,
			-1, 2, 3, -4, -5, -6, -7, -8,
		},
		[16]int8{
			11, 12, 13, 14, 15, 16, 17, 18,
			100, 125, 126, -128, -124, 127, -2, -3,
		},
	)
	expect := [16]int8{
		12, 14, 16, 18, 20, 22, 24, 26,
		99, 127, 127, -128, -128, 121, -9, -11,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmSubInt8(t *testing.T) {
	out := EmmSubInt8(
		[16]int8{
			1, 2, 3, 4, 5, 6, 7, 8,
			127, -128, -3, -127, 0, -1, 1, 8,
		},
		[16]int8{
			1, -2, 1, -2, 0, -1, 1, -8,
			-128, -1, 127, -127, -128, 1, -1, 8,
		},
	)
	expect := [16]int8{
		0, 4, 2, 6, 5, 7, 6, 16,
		127, -127, -128, 0, 127, -2, 2, 0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmAddUint8(t *testing.T) {
	out := EmmAddUint8(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			1, 2, 3, 4, 5, 6, 7, 8,
		},
		[16]uint8{
			1, 2, 3, 4, 0, 0, 7, 8,
			100, 254, 255, 251, 0, 1, 127, 128,
		},
	)
	expect := [16]uint8{
		2, 4, 6, 8, 5, 6, 14, 16,
		101, 255, 255, 255, 5, 7, 134, 136,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmSubUint8(t *testing.T) {
	out := EmmSubUint8(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			255, 254, 255, 255, 127, 6, 1, 8,
		},
		[16]uint8{
			1, 0, 255, 0, 4, 0, 8, 8,
			1, 2, 255, 0, 127, 6, 0, 8,
		},
	)
	expect := [16]uint8{
		0, 2, 0, 4, 1, 6, 0, 0,
		254, 252, 0, 255, 0, 0, 1, 0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmAvgUint8(t *testing.T) {
	out := EmmAvgUint8(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			255, 250, 255, 255, 127, 6, 1, 8,
		},
		[16]uint8{
			9, 2, 4, 16, 127, 121, 121, 16,
			255, 5, 0, 127, 127, 12, 254, 255,
		},
	)
	expect := [16]uint8{
		5, 2, 4, 10, 66, 64, 64, 12,
		255, 128, 128, 191, 127, 9, 128, 132,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmMaxUint8(t *testing.T) {
	out := EmmMaxUint8(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			255, 250, 0, 255, 127, 6, 7, 8,
		},
		[16]uint8{
			0, 1, 10, 255, 254, 0, 7, 8,
			255, 251, 255, 1, 128, 7, 6, 9,
		},
	)
	expect := [16]uint8{
		1, 2, 10, 255, 254, 6, 7, 8,
		255, 251, 255, 255, 128, 7, 7, 9,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmMinUint8(t *testing.T) {
	out := EmmMinUint8(
		[16]uint8{
			1, 2, 3, 4, 5, 6, 7, 8,
			255, 250, 0, 255, 127, 6, 7, 8,
		},
		[16]uint8{
			0, 1, 10, 255, 254, 0, 7, 8,
			255, 251, 255, 1, 128, 7, 6, 9,
		},
	)
	expect := [16]uint8{
		0, 1, 3, 4, 5, 0, 7, 8,
		255, 250, 0, 1, 127, 6, 6, 8,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmAddInt16(t *testing.T) {
	out := EmmAddInt16(
		[8]int16{1, 2, 32767, -32768, 1, 1, -32768, 0},
		[8]int16{1, 3, 1, 1, -32768, 32767, -1, -1},
	)
	expect := [8]int16{
		2, 5, 32767, -32767, -32767, 32767, -32768, -1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmSubInt16(t *testing.T) {
	out := EmmSubInt16(
		[8]int16{1, 2, 32767, -32768, 1, 1, -32768, 0},
		[8]int16{1, 3, 1, 1, -32768, 32767, -1, -1},
	)
	expect := [8]int16{
		0, -1, 32766, -32768, 32767, -32766, -32767, 1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmMaxInt16(t *testing.T) {
	out := EmmMaxInt16(
		[8]int16{1, 2, 32767, -32768, 1, 1, -32768, 0},
		[8]int16{1, 3, 1, 1, -32768, 32767, -1, -1},
	)
	expect := [8]int16{
		1, 3, 32767, 1, 1, 32767, -1, 0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmMinInt16(t *testing.T) {
	out := EmmMinInt16(
		[8]int16{1, 2, 32767, -32768, 1, 1, -32768, 0},
		[8]int16{1, 3, 1, 1, -32768, 32767, -1, -1},
	)
	expect := [8]int16{
		1, 2, 1, -32768, -32768, 1, -32768, -1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmMAddInt16(t *testing.T) {
	out := EmmMAddInt16(
		[8]int16{1, 2, 3, -1, 32767, -32768, -1, -1},
		[8]int16{1, 5, 2, 0, 1, -1, 0, -1},
	)
	expect := [4]int32{
		11, 6, 65535, 1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmAddUint16(t *testing.T) {
	out := EmmAddUint16(
		[8]uint16{1, 2, 32767, 65535, 1, 65535, 0, 0},
		[8]uint16{1, 3, 1, 1, 65535, 65535, 0, 1},
	)
	expect := [8]uint16{
		2, 5, 32768, 65535, 65535, 65535, 0, 1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmSubUint16(t *testing.T) {
	out := EmmSubUint16(
		[8]uint16{0, 1, 1, 0, 65535, 65535, 1, 0},
		[8]uint16{1, 2, 32767, 65535, 1, 65535, 0, 1},
	)
	expect := [8]uint16{
		0, 0, 0, 0, 65534, 0, 1, 0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmAvgUint16(t *testing.T) {
	out := EmmAvgUint16(
		[8]uint16{1, 2, 32767, 65535, 1, 65535, 0, 0},
		[8]uint16{1, 3, 1, 1, 65535, 65535, 0, 1},
	)
	expect := [8]uint16{
		1, 3, 16384, 32768, 32768, 65535, 0, 1,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmAddFloat64(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := EmmAddFloat64(a, b)

	expect := [2]float64{-101.1, 25.2}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmSubFloat64(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := EmmSubFloat64(a, b)

	expect := [2]float64{98.9, -14.8}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmMulFloat64(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := EmmMulFloat64(a, b)

	expect := [2]float64{110.00000000000001, 104.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmDivFloat64(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := EmmDivFloat64(a, b)

	expect := [2]float64{0.011000000000000001, 0.26}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmSqrtFloat64(t *testing.T) {
	a := [2]float64{2.0, 4.0}
	out := EmmSqrtFloat64(a)

	expect := [2]float64{1.4142135623730951, 2}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmMinFloat64(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := EmmMinFloat64(a, b)

	expect := [2]float64{-100.0, 5.2}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmMaxFloat64(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := EmmMaxFloat64(a, b)

	expect := [2]float64{-1.1, 20.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmAndFloat64(t *testing.T) {
	ua := [2]uint64{
		0b0_00000000001_0000000000000000000000000000000000000000000000000000,
		0b0_00011000001_0000000000000000000000000000000000000000000000000000,
	}
	ub := [2]uint64{
		0b0_00000000011_0000000000000000000000000000000000000000000000000000,
		0b0_00010000011_0000000000000000000000000000000000000000000000000000,
	}
	ue := [2]uint64{
		0b0_00000000001_0000000000000000000000000000000000000000000000000000,
		0b0_00010000001_0000000000000000000000000000000000000000000000000000,
	}

	a := *(*[2]float64)(unsafe.Pointer(&ua[0]))
	b := *(*[2]float64)(unsafe.Pointer(&ub[0]))
	e := *(*[2]float64)(unsafe.Pointer(&ue[0]))

	out := EmmAndFloat64(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestEmmOrFloat64(t *testing.T) {
	ua := [2]uint64{
		0b0_00000000001_0000000000000000000000000000000000000000000000000000,
		0b0_00011000001_0000000000000000000000000000000000000000000000000000,
	}
	ub := [2]uint64{
		0b0_00000000011_0000000000000000000000000000000000000000000000000000,
		0b0_00010000011_0000000000000000000000000000000000000000000000000000,
	}
	ue := [2]uint64{
		0b0_00000000011_0000000000000000000000000000000000000000000000000000,
		0b0_00011000011_0000000000000000000000000000000000000000000000000000,
	}

	a := *(*[2]float64)(unsafe.Pointer(&ua[0]))
	b := *(*[2]float64)(unsafe.Pointer(&ub[0]))
	e := *(*[2]float64)(unsafe.Pointer(&ue[0]))

	out := EmmOrFloat64(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestEmmXorFloat64(t *testing.T) {
	a := [2]float64{0.123, 1.235}
	b := [2]float64{1.0, 0.1}
	out := EmmXorFloat64(a, b)
	expect := [2]float64{1.968, 1.7411192570869989}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestEmmAndNotFloat64(t *testing.T) {
	ua := [2]uint64{
		0b0_00000000001_0000000000000000000000000000000000000000000000000000,
		0b0_00011000001_0000000000000000000000000000000000000000000000000000,
	}
	ub := [2]uint64{
		0b0_00000000011_0000000000000000000000000000000000000000000000000000,
		0b0_00010000011_0000000000000000000000000000000000000000000000000000,
	}
	ue := [2]uint64{
		0b0_00000000010_0000000000000000000000000000000000000000000000000000,
		0b0_00000000010_0000000000000000000000000000000000000000000000000000,
	}

	a := *(*[2]float64)(unsafe.Pointer(&ua[0]))
	b := *(*[2]float64)(unsafe.Pointer(&ub[0]))
	e := *(*[2]float64)(unsafe.Pointer(&ue[0]))

	out := EmmAndNotFloat64(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}
