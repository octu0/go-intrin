//go:build fma

package x86

import (
	"reflect"
	"testing"
)

func TestFloat64FMAdd(t *testing.T) {
	a := [2]float64{10.0, 2.0}
	b := [2]float64{2.0, -2.0}
	c := [2]float64{5.0, 3.0}
	out := Float64FMAdd(a, b, c)

	// fmadd(a, b, c) = (a * b) + c
	expect := [2]float64{25.0, -1.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64FMSub(t *testing.T) {
	a := [2]float64{10.0, 2.0}
	b := [2]float64{2.0, -2.0}
	c := [2]float64{5.0, 3.0}
	out := Float64FMSub(a, b, c)

	// fmadd(a, b, c) = (a * b) + c
	expect := [2]float64{15.0, -7.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64FMAdd4(t *testing.T) {
	a := [4]float64{10.0, 2.0, 5.0, -1.0}
	b := [4]float64{2.0, -2.0, 10.0, 8.0}
	c := [4]float64{5.0, 3.0, 2.0, 1.0}
	out := Float64FMAdd4(a, b, c)

	// fmadd(a, b, c) = (a * b) + c
	expect := [4]float64{
		25.0, -1.0, 52.0, -7.0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat64FMSub4(t *testing.T) {
	a := [4]float64{10.0, 2.0, 5.0, -1.0}
	b := [4]float64{2.0, -2.0, 10.0, 8.0}
	c := [4]float64{5.0, 3.0, 2.0, 1.0}
	out := Float64FMSub4(a, b, c)

	// fmadd(a, b, c) = (a * b) + c
	expect := [4]float64{
		15.0, -7.0, 48.0, -9.0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}
