//go:build fma

package x86

import (
	"reflect"
	"testing"
)

func TestFloat32FMAdd(t *testing.T) {
	a := [4]float32{10.0, 2.0, 5.0, -1.0}
	b := [4]float32{2.0, -2.0, 10.0, 8.0}
	c := [4]float32{5.0, 3.0, 2.0, 1.0}
	out := Float32FMAdd(a, b, c)

	// fmadd(a, b, c) = (a * b) + c
	expect := [4]float32{
		25.0, -1.0, 52.0, -7.0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32FMAdd8(t *testing.T) {
	a := [8]float32{10.0, 2.0, 5.0, -1.0, 3.0, 3.0, 300.0, 100.0}
	b := [8]float32{2.0, -2.0, 10.0, 8.0, 3.0, 1.0, 1.0, 1.0}
	c := [8]float32{5.0, 3.0, 2.0, 1.0, 0.5, -1.5, -25.0, -25.0}
	out := Float32FMAdd8(a, b, c)

	// fmadd(a, b, c) = (a * b) + c
	expect := [8]float32{
		25.0, -1.0, 52.0, -7.0, 9.5, 1.5, 275.0, 75.0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}
