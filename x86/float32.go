package x86

type Float32Filter struct {
	Base [4]float32
}

func (f *Float32Filter) Add(values ...float32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	res := XmmBulkAdd(f.Base, aligned, len(aligned))
	return res[:initSize]
}

func (f *Float32Filter) Sub(values ...float32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	res := XmmBulkSub(f.Base, aligned, len(aligned))
	return res[:initSize]
}

func (f *Float32Filter) Mul(values ...float32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	res := XmmBulkMul(f.Base, aligned, len(aligned))
	return res[:initSize]
}

func (f *Float32Filter) Div(values ...float32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	res := XmmBulkDiv(f.Base, aligned, len(aligned))
	return res[:initSize]
}

func Float32Add(a, b [4]float32) [4]float32 {
	return XmmAdd(a, b)
}

func Float32Sub(a, b [4]float32) [4]float32 {
	return XmmSub(a, b)
}

func Float32Mul(a, b [4]float32) [4]float32 {
	return XmmMul(a, b)
}

func Float32Div(a, b [4]float32) [4]float32 {
	return XmmDiv(a, b)
}

func Float32Sqrt(a [4]float32) [4]float32 {
	return XmmSqrt(a)
}

func Float32RSqrt(a [4]float32) [4]float32 {
	return XmmRSqrt(a)
}

func Float32Rcp(a [4]float32) [4]float32 {
	return XmmRcp(a)
}

func Float32Min(a, b [4]float32) [4]float32 {
	return XmmMin(a, b)
}

func Float32Max(a, b [4]float32) [4]float32 {
	return XmmMax(a, b)
}

func Float32And(a, b [4]float32) [4]float32 {
	return XmmAnd(a, b)
}

func Float32Or(a, b [4]float32) [4]float32 {
	return XmmOr(a, b)
}

func Float32Xor(a, b [4]float32) [4]float32 {
	return XmmXor(a, b)
}

func Float32AndNot(a, b [4]float32) [4]float32 {
	return XmmAndNot(a, b)
}

func Int8ToFloat32(values ...int8) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := XmmBulkConvertInt8ToFloat32(aligned, len(aligned))
	return converted[:initSize]
}

func Uint8ToFloat32(values ...uint8) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := XmmBulkConvertUint8ToFloat32(aligned, len(aligned))
	return converted[:initSize]
}

func Float32ToInt8(values ...float32) []int8 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := XmmBulkConvertFloat32ToInt8(aligned, len(aligned))
	return converted[:initSize]
}

func Float32ToUint8(values ...float32) []uint8 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := XmmBulkConvertFloat32ToUint8(aligned, len(aligned))
	return converted[:initSize]
}

func Float32Tile4Sum(values ...float32) []float32 {
	aligned := alignSlice(values, 16)
	tiled := XmmTile4x4Sum(aligned, len(aligned))
	return tiled
}

func Float32Sum(values ...float32) float32 {
	if len(values) < 256 {
		return float32SumNative(values)
	}

	aligned := alignSlice(values, 4)
	return XmmBulkSum(aligned, len(aligned))
}

func float32SumNative(values []float32) float32 {
	sum := float32(0.0)
	for _, v := range values {
		sum += v
	}
	return sum
}

func convertFloat32[T number](values []T) []float32 {
	out := make([]float32, len(values))
	for i, v := range values {
		out[i] = float32(v)
	}
	return out
}

func convertUint8[T number](values []T) []uint8 {
	out := make([]uint8, len(values))
	for i, v := range values {
		out[i] = uint8(v)
	}
	return out
}
