package iterator

type OfBool = Iterator[bool]
type OfByte = Iterator[byte]
type OfFloat32 = Iterator[float32]
type OfFloat64 = Iterator[float64]
type OfInt = Iterator[int]
type OfInt32 = Iterator[int32]
type OfInt64 = Iterator[int64]
type OfRune = Iterator[rune]
type OfString = Iterator[string]

type OfBoolSlice = Iterator[[]bool]
type OfByteSlice = Iterator[[]byte]
type OfFloat32Slice = Iterator[[]float32]
type OfFloat64Slice = Iterator[[]float64]
type OfIntSlice = Iterator[[]int]
type OfInt32Slice = Iterator[[]int32]
type OfInt64Slice = Iterator[[]int64]
type OfRuneSlice = Iterator[[]rune]
type OfStringSlice = Iterator[[]string]

type Iterator[T any] func() (T, Iterator[T])

func FromSlice[T any](ts []T) Iterator[T] {
	if len(ts) == 0 {
		return nil
	}

	i := 0

	var next Iterator[T]
	next = func() (T, Iterator[T]) {
		if i == len(ts)-1 {
			return ts[i], nil
		}

		t := ts[i]
		i += 1
		return t, next
	}

	return next
}
