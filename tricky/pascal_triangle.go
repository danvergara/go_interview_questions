package tricky

import "fmt"

// Pascal prints out a Pascal's triangle, given the number of desired rows.
func Pascal(n int) {
	arr := []int{1}
	temp := []int{}

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
