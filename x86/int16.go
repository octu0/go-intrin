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
