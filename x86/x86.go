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
