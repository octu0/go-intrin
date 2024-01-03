package x86

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"os"
	"reflect"
	"testing"

	_ "embed"
)

var (
	//go:embed testdata/src.png
	pngImg []byte
)

func pngToRGBA(data []byte) (*image.NRGBA, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	if i, ok := img.(*image.NRGBA); ok {
		return i, nil
	}

	b := img.Bounds()
	nrgba := image.NewNRGBA(b)
	for y := b.Min.Y; y < b.Max.Y; y += 1 {
		for x := b.Min.X; x < b.Max.X; x += 1 {
			c := color.NRGBAModel.Convert(img.At(x, y))
			nrgba.Set(x, y, c)
		}
	}
	return nrgba, nil
}

func goGrayscaleFloat32(src *image.NRGBA) *image.NRGBA {
	b := src.Bounds()
	w, h := b.Dx(), b.Dy()

	out := image.NewNRGBA(b)
	for y := 0; y < h; y += 1 {
		for x := 0; x < w; x += 1 {
			c := src.NRGBAAt(x, y)
			// BT.709
			gray := byte((0.2126 * float32(c.R)) + (0.7152 * float32(c.G)) + (0.0722 * float32(c.B)))
			out.SetNRGBA(x, y, color.NRGBA{
				R: gray,
				G: gray,
				B: gray,
				A: 0xff,
			})
		}
	}
	return out
}

func xmmGrayscaleFloat32(src *image.NRGBA) *image.NRGBA {
	data := Uint8ToFloat32(src.Pix...)

	/*
			f := Float32Filter{Base: [4]float32{0.2126, 0.7152, 0.0722, 0.0}}
			filtered := f.Mul(data...)
			pix := make([]byte, len(data))
			for i := 0; i < len(data); i += 16 {
				rg := Float32Add(
					[4]float32{filtered[i+0+0], filtered[i+4+0], filtered[i+8+0], filtered[i+12+0]}, // R
					[4]float32{filtered[i+0+1], filtered[i+4+1], filtered[i+8+1], filtered[i+12+1]}, // G
				)
				ba := Float32Add(
					[4]float32{filtered[i+0+2], filtered[i+4+2], filtered[i+8+2], filtered[i+12+2]}, // B
					[4]float32{filtered[i+0+3], filtered[i+4+3], filtered[i+8+3], filtered[i+12+3]}, // A
				)
		    out := Float32Add(rg, ba)
		    data := Float32ToUint8(out[0], out[1], out[2], out[3])
				pix[i+0+0] = data[0]  // R
				pix[i+0+1] = data[0]  // G
				pix[i+0+2] = data[0]  // B
				pix[i+0+3] = 0xff     // A
				pix[i+4+0] = data[1]  // R
				pix[i+4+1] = data[1]  // G
				pix[i+4+2] = data[1]  // B
				pix[i+4+3] = 0xff     // A
				pix[i+8+0] = data[2]  // R
				pix[i+8+1] = data[2]  // G
				pix[i+8+2] = data[2]  // B
				pix[i+8+3] = 0xff     // A
				pix[i+12+0] = data[3] // R
				pix[i+12+1] = data[3] // G
				pix[i+12+2] = data[3] // B
				pix[i+12+3] = 0xff    // A
			}
	*/

	f := Float32Filter{Base: [4]float32{0.2126, 0.7152, 0.0722, 0.0}}
	filtered := f.Mul(data...)
	tiledGray := Float32Tile4Sum(filtered...)
	pix := make([]byte, 0, len(src.Pix))
	for _, gray := range Float32ToUint8(tiledGray...) {
		pix = append(pix, gray) // R
		pix = append(pix, gray) // G
		pix = append(pix, gray) // B
		pix = append(pix, 0xff) // A
	}

	return &image.NRGBA{
		Pix:    pix,
		Stride: src.Stride,
		Rect:   src.Rect,
	}
}

func saveImage(img *image.NRGBA) (string, error) {
	out, err := os.CreateTemp("/tmp", "out*.png")
	if err != nil {
		return "", err
	}
	defer out.Close()

	if err := png.Encode(out, img); err != nil {
		return "", err
	}
	return out.Name(), nil
}

func BenchmarkFloat32(b *testing.B) {
	b.Run("sum/go", func(tb *testing.B) {
		v := make([]float32, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = float32(i)
		}
		tb.ResetTimer()
		_ = float32SumNative(v)
	})
	b.Run("sum/xmm", func(tb *testing.B) {
		v := make([]float32, tb.N)
		for i := 0; i < tb.N; i += 1 {
			v[i] = float32(i)
		}
		tb.ResetTimer()
		_ = Float32Sum(v...)
	})
}

func BenchmarkGrayscaleFloat32(b *testing.B) {
	b.Run("go", func(tb *testing.B) {
		img, err := pngToRGBA(pngImg)
		if err != nil {
			tb.Fatalf("%+v", err)
		}
		/*
		   a := goGrayscaleFloat32(img)
		   out, err := saveImage(a)
		   if err != nil {
		     tb.Fatalf("%+v", err)
		   }
		   println("go grayscale =", out)
		*/
		tb.ResetTimer()
		for i := 0; i < tb.N; i += 1 {
			_ = goGrayscaleFloat32(img)
		}
	})
	b.Run("xmm", func(tb *testing.B) {
		img, err := pngToRGBA(pngImg)
		if err != nil {
			tb.Fatalf("%+v", err)
		}
		/*
		   a := xmmGrayscaleFloat32(img)
		   out, err := saveImage(a)
		   if err != nil {
		     tb.Fatalf("%+v", err)
		   }
		   println("xmm grayscale =", out)
		*/
		tb.ResetTimer()
		for i := 0; i < tb.N; i += 1 {
			_ = xmmGrayscaleFloat32(img)
		}
	})

}

func TestFloat32Filter(t *testing.T) {
	t.Run("add", func(tt *testing.T) {
		f := Float32Filter{Base: [4]float32{128.0, 64.0, 32.0, 16.0}}
		out := f.Add(
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			6.0,
		)
		expect := []float32{
			129.0,
			66.0,
			35.0,
			20.0,
			133.0,
			70.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("sub", func(tt *testing.T) {
		f := Float32Filter{Base: [4]float32{128.0, 64.0, 32.0, 16.0}}
		out := f.Sub(
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			6.0,
		)
		expect := []float32{
			-127.0,
			-62.0,
			-29.0,
			-12.0,
			-123.0,
			-58.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("mul", func(tt *testing.T) {
		f := Float32Filter{Base: [4]float32{2.0, 3.0, 4.0, 8.0}}
		out := f.Mul(
			1.0,
			2.0,
			3.0,
			4.0,
			5.0,
			6.0,
		)
		expect := []float32{
			2.0,
			6.0,
			12.0,
			32.0,
			10.0,
			18.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("div", func(tt *testing.T) {
		f := Float32Filter{Base: [4]float32{2.0, 3.0, 4.0, 8.0}}
		out := f.Div(
			1.0,
			2.1,
			3.0,
			4.0,
			5.0,
			6.0,
		)
		expect := []float32{
			0.5,
			0.7,
			0.75,
			0.5,
			2.5,
			2.0,
		}
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
}

func TestInt8ToFloat32(t *testing.T) {
	out := Int8ToFloat32(1, 2, 127, -128, -1, 0)
	expect := []float32{1.0, 2.0, 127.0, -128.0, -1.0, 0.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestUint8ToFloat32(t *testing.T) {
	out := Uint8ToFloat32(1, 2, 127, 255, 128, 0)
	expect := []float32{1.0, 2.0, 127.0, 255.0, 128.0, 0.0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32ToInt8(t *testing.T) {
	out := Float32ToInt8(1.0, 2.0, 127.0, 255.0, -128.0, -1)
	expect := []int8{1, 2, 127, 127, -128, -1}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32ToUint8(t *testing.T) {
	out := Float32ToUint8(1.0, 2.0, 127.0, 255.0, -128.0, -1, -2, 0)
	//expect := []uint8{1, 2, 127, 255, 128, 255, 254, 0}
	expect := []uint8{1, 2, 127, 127, 128, 255, 254, 0}
	if reflect.DeepEqual(expect, out) != true {
		t.Errorf("expect %v <> actual %v", expect, out)
	}
}

func TestFloat32Sum(t *testing.T) {
	t.Run("size=4", func(tt *testing.T) {
		out := Float32Sum(1.0, 2.0, 3.0, 4.0)
		expect := float32(10.0)
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("odd", func(tt *testing.T) {
		out := Float32Sum(1.0, 2.0, 3.0, 4.0, 5.0)
		expect := float32(15.0)
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
	t.Run("large", func(tt *testing.T) {
		N := 1023
		v := make([]float32, N)
		for i := 0; i < N; i += 1 {
			v[i] = float32(i)
		}

		out := Float32Sum(v...)
		expect := float32SumNative(v)
		if reflect.DeepEqual(expect, out) != true {
			tt.Errorf("expect %v <> actual %v", expect, out)
		}
	})
}
