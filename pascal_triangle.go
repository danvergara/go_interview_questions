package main

import "fmt"

func printPascal(n int) {
	arr := []int{1}
	temp := []int{}

	fmt.Printf("pascal's triangle of %d, rows...\n", n)

	for i := 0; i < n; i++ {
		fmt.Printf("rows %d : ", i+1)
		for _, j := range arr {
			fmt.Printf("%d ", j)
		}

		fmt.Println()
		temp = append(temp, 1)

		for j := 0; j < len(arr)-1; j++ {
			temp = append(temp, arr[j]+arr[j+1])
		}

		temp = append(temp, 1)
		arr = temp
		temp = []int{}
	}
}
