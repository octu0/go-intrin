package x86

type number interface {
	~int8 | ~int16 | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func alignSlice[T number](values []T, size int) []T {
	if (len(values) % size) == 0 {
		return values
	}

	remain := size - (len(values) % size)
	if len(values) < size {
		remain = size - len(values)
	}
	for i := 0; i < remain; i += 1 {
		values = append(values, 0)
	}
	return values
}

func convertInt8[T number](values []T) []int8 {
	out := make([]int8, len(values))
	for i, v := range values {
		out[i] = int8(v)
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

func convertInt16[T number](values []T) []int16 {
	out := make([]int16, len(values))
	for i, v := range values {
		out[i] = int16(v)
	}
	return out
}

func convertUint16[T number](values []T) []uint16 {
	out := make([]uint16, len(values))
	for i, v := range values {
		out[i] = uint16(v)
	}
	return out
}

func convertInt32[T number](values []T) []int32 {
	out := make([]int32, len(values))
	for i, v := range values {
		out[i] = int32(v)
	}
	return out
}

func convertUint32[T number](values []T) []uint32 {
	out := make([]uint32, len(values))
	for i, v := range values {
		out[i] = uint32(v)
	}
	return out
}

func convertFloat32[T number](values []T) []float32 {
	out := make([]float32, len(values))
	for i, v := range values {
		out[i] = float32(v)
	}
	return out
}

func convertFloat64[T number](values []T) []float64 {
	out := make([]float64, len(values))
	for i, v := range values {
		out[i] = float64(v)
	}
	return out
}
