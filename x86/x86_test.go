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

func xmmGrayscaleFloat32_16(src *image.NRGBA) *image.NRGBA {
	data := Uint8ToFloat32(src.Pix...)

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

	return &image.NRGBA{
		Pix:    pix,
		Stride: src.Stride,
		Rect:   src.Rect,
	}
}

func xmmGrayscaleFloat32_tile(src *image.NRGBA) *image.NRGBA {
	data := Uint8ToFloat32(src.Pix...)
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

func xmmGrayscaleFloat32(src *image.NRGBA) *image.NRGBA {
	initSize := len(src.Pix)
	aligned := alignSlice(src.Pix, 16)
	pix := xmmRGBAGrayscale(aligned, len(aligned))
	return &image.NRGBA{
		Pix:    pix[:initSize],
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

func BenchmarkGrayscaleFloat32(b *testing.B) {
	b.Run("go", func(tb *testing.B) {
		img, err := pngToRGBA(pngImg)
		if err != nil {
			tb.Fatalf("%+v", err)
		}
		tb.ResetTimer()
		for i := 0; i < tb.N; i += 1 {
			_ = goGrayscaleFloat32(img)
		}
	})
	b.Run("simd/small", func(tb *testing.B) {
		img, err := pngToRGBA(pngImg)
		if err != nil {
			tb.Fatalf("%+v", err)
		}
		tb.ResetTimer()
		for i := 0; i < tb.N; i += 1 {
			_ = xmmGrayscaleFloat32_16(img)
		}
	})
	b.Run("simd/medium", func(tb *testing.B) {
		img, err := pngToRGBA(pngImg)
		if err != nil {
			tb.Fatalf("%+v", err)
		}
		tb.ResetTimer()
		for i := 0; i < tb.N; i += 1 {
			_ = xmmGrayscaleFloat32_tile(img)
		}
	})
	b.Run("simd/full", func(tb *testing.B) {
		img, err := pngToRGBA(pngImg)
		if err != nil {
			tb.Fatalf("%+v", err)
		}
		tb.ResetTimer()
		for i := 0; i < tb.N; i += 1 {
			_ = xmmGrayscaleFloat32(img)
		}
	})
}

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

func init() {
	if false {
		img, err := pngToRGBA(pngImg)
		if err != nil {
			panic(err)
		}
		a := goGrayscaleFloat32(img)
		out, err := saveImage(a)
		if err != nil {
			panic(err)
		}
		println("go grayscale =", out)
	}
	if false {
		img, err := pngToRGBA(pngImg)
		if err != nil {
			panic(err)
		}
		a := xmmGrayscaleFloat32(img)
		out, err := saveImage(a)
		if err != nil {
			panic(err)
		}
		println("xmm grayscale =", out)
	}
}
