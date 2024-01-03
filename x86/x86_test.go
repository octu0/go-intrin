package x86

import (
	"reflect"
	"testing"
)

func TestAlignSlice(t *testing.T) {
	t.Run("size=4/align=4", func(tt *testing.T) {
		in := []float32{1.0, 2.0, 3.0, 4.0}
		out := alignSlice(in, 4)

		expect := []float32{1.0, 2.0, 3.0, 4.0}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("odd/align=4", func(tt *testing.T) {
		in := []float32{1.0, 2.0, 3.0, 4.0, 5.0}
		out := alignSlice(in, 4)

		expect := []float32{1.0, 2.0, 3.0, 4.0, 5.0, 0.0, 0.0, 0.0}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("less/align=4", func(tt *testing.T) {
		in := []int8{1, 2}
		out := alignSlice(in, 4)

		expect := []int8{1, 2, 0, 0}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
}
