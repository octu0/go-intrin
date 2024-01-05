package x86

func Int32ToFloat64(values ...int32) []float64 {
	if 1024 < len(values) {
		return int32ToFloat64M256(values)
	}
	return int32ToFloat64M128(values)
}

func int32ToFloat64M128(values []int32) []float64 {
	initSize := len(values)
	aligned := alignSlice(values, 2)
	converted := emmBulkConvertInt32ToFloat64(aligned, len(aligned))
	return converted[:initSize]
}

func int32ToFloat64M256(values []int32) []float64 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := immBulkConvertInt32ToFloat64(aligned, len(aligned))
	return converted[:initSize]
}

func Float64ToInt32(values ...float64) []int32 {
	if 1024 < len(values) {
		return float64ToInt32M256(values)
	}
	return float64ToInt32M128(values)
}

func float64ToInt32M128(values []float64) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 2)
	converted := emmBulkConvertFloat64ToInt32(aligned, len(aligned))
	return converted[:initSize]
}

func float64ToInt32M256(values []float64) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := immBulkConvertFloat64ToInt32(aligned, len(aligned))
	return converted[:initSize]
}

func Int32ToFloat32(values ...int32) []float32 {
	if 1024 < len(values) {
		return int32ToFloat32M256(values)
	}
	return int32ToFloat32M128(values)
}

func int32ToFloat32M128(values []int32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := emmBulkConvertInt32ToFloat32(aligned, len(aligned))
	return converted[:initSize]
}

func int32ToFloat32M256(values []int32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 8)
	converted := immBulkConvertInt32ToFloat32(aligned, len(aligned))
	return converted[:initSize]
}

func Float32ToInt32(values ...float32) []int32 {
	if 1024 < len(values) {
		return float32ToInt32M256(values)
	}
	return float32ToInt32M128(values)
}

func float32ToInt32M128(values []float32) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := emmBulkConvertFloat32ToInt32(aligned, len(aligned))
	return converted[:initSize]
}

func float32ToInt32M256(values []float32) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 8)
	converted := immBulkConvertFloat32ToInt32(aligned, len(aligned))
	return converted[:initSize]
}
