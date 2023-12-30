package x86

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestXmmAdd(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := XmmAdd(a, b)

	expect := [4]float32{-101.1, 25.2, 0.0, -15.3}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestXmmSub(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := XmmSub(a, b)

	expect := [4]float32{98.9, -14.8, 0.0, -5.3}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestXmmMul(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := XmmMul(a, b)

	expect := [4]float32{110.0, 104.0, 0.0, 51.5}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestXmmDiv(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.4, -5.0}
	out := XmmDiv(a, b)

	expect := [4]float32{0.011, 0.26, 0.0, 2.06}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestXmmSqrt(t *testing.T) {
	a := [4]float32{2.0, 4.0, 13.0, 16.0}
	out := XmmSqrt(a)

	expect := [4]float32{1.4142135, 2, 3.6055512, 4}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestXmmRSqrt(t *testing.T) {
	a := [4]float32{2.0, 4.0, 13.0, 16.0}
	out := XmmRSqrt(a)

	expect := [4]float32{0.7069092, 0.49987793, 0.2772827, 0.24993896}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestXmmRcp(t *testing.T) {
	a := [4]float32{1.0, 4.0, 13.0, 100}
	out := XmmRcp(a)

	expect := [4]float32{0.99975586, 0.24993896, 0.0769043, 0.0099983215}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestXmmMin(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := XmmMin(a, b)

	expect := [4]float32{-100.0, 5.2, 0.0, -10.3}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestXmmMax(t *testing.T) {
	a := [4]float32{-1.1, 5.2, 0.0, -10.3}
	b := [4]float32{-100.0, 20.0, 0.0, -5.0}
	out := XmmMax(a, b)

	expect := [4]float32{-1.1, 20.0, 0.0, -5.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestXmmAnd(t *testing.T) {
	ua := [4]uint32{
		0b0_00000001_0000000000000000000000,
		0b0_11000001_0000000000000000000000,
		0,
		0b0_10000001_1001000000000000000001,
	}
	ub := [4]uint32{
		0b0_00000011_0000000000000000000000,
		0b0_10000011_0000000000000000000000,
		0,
		0b0_10000000_1111000000000000000000,
	}
	ue := [4]uint32{
		0b0_00000001_0000000000000000000000,
		0b0_10000001_0000000000000000000000,
		0,
		0b0_10000000_1001000000000000000000,
	}

	a := *(*[4]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[4]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[4]float32)(unsafe.Pointer(&ue[0]))

	out := XmmAnd(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestXmmOr(t *testing.T) {
	ua := [4]uint32{
		0b0_00000001_0000000000000000000000,
		0b0_11000001_0000000000000000000000,
		0,
		0b0_10000001_1001000000000000000001,
	}
	ub := [4]uint32{
		0b0_00000011_0000000000000000000000,
		0b0_10000011_0000000000000000000000,
		0,
		0b0_10000000_1111000000000000000000,
	}
	ue := [4]uint32{
		0b0_00000011_0000000000000000000000,
		0b0_11000011_0000000000000000000000,
		0,
		0b0_10000001_1111000000000000000001,
	}

	a := *(*[4]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[4]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[4]float32)(unsafe.Pointer(&ue[0]))

	out := XmmOr(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestXmmXor(t *testing.T) {
	ua := [4]uint32{
		0b0_00000001_0000000000000000000000,
		0b0_11000001_0000000000000000000000,
		0,
		0b0_10000001_1001000000000000000001,
	}
	ub := [4]uint32{
		0b0_00000011_0000000000000000000000,
		0b0_10000011_0000000000000000000000,
		0,
		0b0_10000000_1111000000000000000000,
	}
	ue := [4]uint32{
		0b0_00000010_0000000000000000000000,
		0b0_01000010_0000000000000000000000,
		0,
		0b0_00000001_0110000000000000000001,
	}

	a := *(*[4]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[4]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[4]float32)(unsafe.Pointer(&ue[0]))

	out := XmmXor(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestXmmAndNot(t *testing.T) {
	ua := [4]uint32{
		0b0_00000001_0000000000000000000000,
		0b0_11000001_0000000000000000000000,
		0,
		0b0_10000001_1001000000000000000001,
	}
	ub := [4]uint32{
		0b0_00000011_0000000000000000000000,
		0b0_10000011_0000000000000000000000,
		0,
		0b0_10000000_1111000000000000000000,
	}
	ue := [4]uint32{
		0b0_00000010_0000000000000000000000,
		0b0_00000010_0000000000000000000000,
		0,
		0b0_00000000_0110000000000000000000,
	}

	a := *(*[4]float32)(unsafe.Pointer(&ua[0]))
	b := *(*[4]float32)(unsafe.Pointer(&ub[0]))
	e := *(*[4]float32)(unsafe.Pointer(&ue[0]))

	out := XmmAndNot(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}
