package x86

import (
	"reflect"
	"testing"
)

func TestInt32ToFloat32(t *testing.T) {
	out := Int32ToFloat32(100, -100, 32767, 65535, 123456)
	expect := []float32{
		100.0, -100.0, 32767.0, 65535.0, 123456.0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32ToInt32(t *testing.T) {
	out := Float32ToInt32(100.0, -100.0, 32767.0, 65535.0, 0.12345)
	expect := []int32{
		100, -100, 32767, 65535, 0,
	}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}
