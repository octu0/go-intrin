package x86

func Int8Add(a, b [16]int8) [16]int8 {
	return EmmAddInt8(a, b)
}

func Int8Sub(a, b [16]int8) [16]int8 {
	return EmmSubInt8(a, b)
}

func Int8Sum(values ...int8) int8 {
	if len(values) < 128 {
		return int8SumNative(values)
	}

	aligned := alignSlice(values, 16)
	return EmmBulkSumInt8(aligned, len(aligned))
}

func int8SumNative(values []int8) int8 {
	sum := int8(0)
	for _, v := range values {
		sum += v
	}
	return sum
}

func Uint8Add(a, b [16]uint8) [16]uint8 {
	return EmmAddUint8(a, b)
}

func Uint8Sub(a, b [16]uint8) [16]uint8 {
	return EmmSubUint8(a, b)
}

func Uint8Avg(a, b [16]uint8) [16]uint8 {
	return EmmAvgUint8(a, b)
}

func Uint8Max(a, b [16]uint8) [16]uint8 {
	return EmmMaxUint8(a, b)
}

func Uint8Min(a, b [16]uint8) [16]uint8 {
	return EmmMinUint8(a, b)
}

func Uint8Sum(values ...uint8) uint8 {
	if len(values) < 128 {
		return uint8SumNative(values)
	}

	aligned := alignSlice(values, 16)
	return EmmBulkSumUint8(aligned, len(aligned))
}

func uint8SumNative(values []uint8) uint8 {
	sum := uint8(0)
	for _, v := range values {
		sum += v
	}
	return sum
}
