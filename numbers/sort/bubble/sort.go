package sort

type Number interface {
	int | int16 | int32 | int64 | float32 | float64
}

// Sort implements the bubble sort algorithm.
// Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in the wrong order. This algorithm is not suitable for large data sets as its average and worst-case time complexity is quite high.
func Sort[N Number](list []N) []N {
	for i := 0; i < len(list); i++ {
		for j := len(list) - 1; j > i; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}

	return list
}
