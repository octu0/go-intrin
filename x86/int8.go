package x86

func Int8Add(a, b [16]int8) [16]int8 {
	return EmmAddInt8(a, b)
}

func Int8Sub(a, b [16]int8) [16]int8 {
	return EmmSubInt8(a, b)
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
