package x86

func Int16Add(a, b [8]int16) [8]int16 {
	return EmmAddInt16(a, b)
}

func Int16Sub(a, b [8]int16) [8]int16 {
	return EmmSubInt16(a, b)
}

func Int16Max(a, b [8]int16) [8]int16 {
	return EmmMaxInt16(a, b)
}

func Int16Min(a, b [8]int16) [8]int16 {
	return EmmMinInt16(a, b)
}

func Int16MAdd(a, b [8]int16) [4]int32 {
	return EmmMAddInt16(a, b)
}

func Int16Sum(values ...int16) int16 {
	if len(values) < 256 {
		return int16SumNative(values)
	}

	aligned := alignSlice(values, 8)
	return EmmBulkSumInt16(aligned, len(aligned))
}

func int16SumNative(values []int16) int16 {
	sum := int16(0)
	for _, v := range values {
		sum += v
	}
	return sum
}

func Uint16Add(a, b [8]uint16) [8]uint16 {
	return EmmAddUint16(a, b)
}

func Uint16Sub(a, b [8]uint16) [8]uint16 {
	return EmmSubUint16(a, b)
}

func Uint16Avg(a, b [8]uint16) [8]uint16 {
	return EmmAvgUint16(a, b)
}

func Uint16Sum(values ...uint16) uint16 {
	if len(values) < 256 {
		return uint16SumNative(values)
	}

	aligned := alignSlice(values, 8)
	return EmmBulkSumUint16(aligned, len(aligned))
}

func uint16SumNative(values []uint16) uint16 {
	sum := uint16(0)
	for _, v := range values {
		sum += v
	}
	return sum
}
