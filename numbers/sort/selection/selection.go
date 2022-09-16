package selection

type Number interface {
	int | int16 | int32 | int64 | float32 | float64
}

// Sort implements the selection sort algorithm.
// It sorts an array by repeatedly finding the minimum element (considering ascending order) from the unsorted part and putting it at the beginning.
// The algorithm maintains two subarrays in a given array.
// The subarray which already sorted.
// The remaining subarray was unsorted.
func Sort[N Number](list []N) []N {
	for i := 0; i < len(list)-1; i++ {
		minimum := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minimum] {
				minimum = j
			}
		}
		if minimum != i {
			list[i], list[minimum] = list[minimum], list[i]
		}
	}

	return list
}
