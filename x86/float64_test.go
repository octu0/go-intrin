package x86

import (
	"reflect"
	"testing"
)

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
