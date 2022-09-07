package main

func insertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		currentNumber := list[i]
		for j := i - 1; j > -1; j-- {
			if list[j] > currentNumber {
				list[j], list[j+1] = list[j+1], list[j]
			} else {
				list[j+1] = currentNumber
				break
			}
		}
	}
	return list
}
