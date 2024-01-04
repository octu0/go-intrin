package x86

func Float64Add(a, b [2]float64) [2]float64 {
	return EmmAddFloat64(a, b)
}

func Float64Sub(a, b [2]float64) [2]float64 {
	return EmmSubFloat64(a, b)
}

func Float64Mul(a, b [2]float64) [2]float64 {
	return EmmMulFloat64(a, b)
}

func Float64Div(a, b [2]float64) [2]float64 {
	return EmmDivFloat64(a, b)
}

func Float64Sqrt(a [2]float64) [2]float64 {
	return EmmSqrtFloat64(a)
}

func Float64Min(a, b [2]float64) [2]float64 {
	return EmmMinFloat64(a, b)
}

func Float64Max(a, b [2]float64) [2]float64 {
	return EmmMaxFloat64(a, b)
}

func Float64And(a, b [2]float64) [2]float64 {
	return EmmAndFloat64(a, b)
}

func Float64Or(a, b [2]float64) [2]float64 {
	return EmmOrFloat64(a, b)
}

func Float64Xor(a, b [2]float64) [2]float64 {
	return EmmXorFloat64(a, b)
}

func Float64AndNot(a, b [2]float64) [2]float64 {
	return EmmAndNotFloat64(a, b)
}

func Int32ToFloat64(values ...int32) []float64 {
	initSize := len(values)
	aligned := alignSlice(values, 2)
	converted := EmmBulkConvertInt32ToFloat64(aligned, len(aligned))
	return converted[:initSize]
}

func Float64ToInt32(values ...float64) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 2)
	converted := EmmBulkConvertFloat64ToInt32(aligned, len(aligned))
	return converted[:initSize]
}
