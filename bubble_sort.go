package main

func bubbleSort(list []int) []int {
	for i := 0; i < len(list); i++ {
		for j := len(list) - 1; j > i; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}

	return list
}
