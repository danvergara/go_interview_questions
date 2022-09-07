package main

func selectionSort(list []int) []int {
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
