package x86

import (
	"reflect"
	"testing"
)

func TestInt32ToFloat32(t *testing.T) {
	t.Run("m128", func(tt *testing.T) {
		out := int32ToFloat32M128([]int32{100, -100, 32767, 65535, 123456})
		expect := []float32{
			100.0, -100.0, 32767.0, 65535.0, 123456.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("m256", func(tt *testing.T) {
		out := int32ToFloat32M256([]int32{100, -100, 32767, 65535, 123456})
		expect := []float32{
			100.0, -100.0, 32767.0, 65535.0, 123456.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
}

func TestFloat32ToInt32(t *testing.T) {
	t.Run("m128", func(tt *testing.T) {
		out := float32ToInt32M128([]float32{100.0, -100.0, 32767.0, 65535.0, 0.12345})
		expect := []int32{
			100, -100, 32767, 65535, 0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("m256", func(tt *testing.T) {
		out := float32ToInt32M256([]float32{100.0, -100.0, 32767.0, 65535.0, 0.12345})
		expect := []int32{
			100, -100, 32767, 65535, 0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
}

func TestInt32ToFloat64(t *testing.T) {
	t.Run("m128", func(tt *testing.T) {
		out := int32ToFloat64M128([]int32{100, -100, 2, 5})
		expect := []float64{100.0, -100.0, 2.0, 5.0}

		if reflect.DeepEqual(expect, out) != true {
			t.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("m256", func(tt *testing.T) {
		out := int32ToFloat64M256([]int32{100, -100, 2, 5})
		expect := []float64{100.0, -100.0, 2.0, 5.0}

		if reflect.DeepEqual(expect, out) != true {
			t.Errorf("expect %v <> actual %v", expect, out)
		}
	})
}

func TestFloat64ToInt32(t *testing.T) {
	t.Run("m128", func(tt *testing.T) {
		out := float64ToInt32M128([]float64{100.0, -100.0, 2.0, 5.0})
		expect := []int32{100, -100, 2, 5}

		if reflect.DeepEqual(expect, out) != true {
			t.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("m256", func(tt *testing.T) {
		out := float64ToInt32M256([]float64{100.0, -100.0, 2.0, 5.0})
		expect := []int32{100, -100, 2, 5}

		if reflect.DeepEqual(expect, out) != true {
			t.Errorf("expect %v <> actual %v", expect, out)
		}
	})
}
