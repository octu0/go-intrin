package x86

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestFloat64Add(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := Float64Add(a, b)

	expect := [2]float64{-101.1, 25.2}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64Sub(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := Float64Sub(a, b)

	expect := [2]float64{98.9, -14.8}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64Mul(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := Float64Mul(a, b)

	expect := [2]float64{110.00000000000001, 104.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64Div(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := Float64Div(a, b)

	expect := [2]float64{0.011000000000000001, 0.26}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64Sqrt(t *testing.T) {
	a := [2]float64{2.0, 4.0}
	out := Float64Sqrt(a)

	expect := [2]float64{1.4142135623730951, 2}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64Min(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := Float64Min(a, b)

	expect := [2]float64{-100.0, 5.2}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64Max(t *testing.T) {
	a := [2]float64{-1.1, 5.2}
	b := [2]float64{-100.0, 20.0}
	out := Float64Max(a, b)

	expect := [2]float64{-1.1, 20.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64And(t *testing.T) {
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

	out := Float64And(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat64Or(t *testing.T) {
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

	out := Float64Or(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestFloat64Xor(t *testing.T) {
	a := [2]float64{0.123, 1.235}
	b := [2]float64{1.0, 0.1}
	out := Float64Xor(a, b)
	expect := [2]float64{1.968, 1.7411192570869989}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64AndNot(t *testing.T) {
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

	out := Float64AndNot(a, b)
	if reflect.DeepEqual(e, out) != true {
		t.Errorf("expect %v <> actual %v", e, out)
	}
}

func TestInt32ToFloat64(t *testing.T) {
	out := Int32ToFloat64(100, -100, 2, 5)
	expect := []float64{100.0, -100.0, 2.0, 5.0}

	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64ToInt32(t *testing.T) {
	out := Float64ToInt32(100.0, -100.0, 2.0, 5.0)
	expect := []int32{100, -100, 2, 5}

	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}
