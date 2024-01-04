package x86

func Int32ToFloat32(values ...int32) []float32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := EmmBulkConvertInt32ToFloat32(aligned, len(aligned))
	return converted[:initSize]
}

func Float32ToInt32(values ...float32) []int32 {
	initSize := len(values)
	aligned := alignSlice(values, 4)
	converted := EmmBulkConvertFloat32ToInt32(aligned, len(aligned))
	return converted[:initSize]
}
