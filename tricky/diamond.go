package tricky

import "fmt"

// Diamond prints out a full diamond pattern on the terminal.
func Diamond(rows int) {
	for i := 1; i < rows+1; i++ {
		for j := 1; j < rows-i+1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Print(" ")
		}
		for l := 1; l < 2*(rows-i); l++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
