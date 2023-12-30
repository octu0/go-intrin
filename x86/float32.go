package x86

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
