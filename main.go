package main

import "fmt"

func main() {
	countCharacters("clcoding")

	palindromeCandidates := "Civic level noon support career"
	fmt.Printf("sentences %s, palindromes: %v\n", palindromeCandidates, palindrome(palindromeCandidates))

	fmt.Printf("birth year: %d \n month year: %d\n birth date: %d\n your age is %d\n", 1992, 11, 16, ageCalculator(1992, 11, 16))

	fmt.Println(anagram("cinema", "iceman"))
	fmt.Println(anagram("cool", "loco"))
	fmt.Println(anagram("men", "women"))
	fmt.Println(anagram("python", "pythno"))

	fmt.Println()
	printPascal(4)

	fmt.Println()
	diamond(7)

	fmt.Println()
	fmt.Printf("sorted list: %v", selectionSort([]int{4, 6, 9, 8, 1, 7, 3}))

	fmt.Println()
	fmt.Printf("sorted list: %v", bubbleSort([]int{7, 1, 8, 2, 9, 4, 6, 5}))

	fmt.Println()
	fmt.Printf("sorted list: %v", insertionSort([]int{3, 7, 2, 8, 4, 1, 9, 5}))

	fmt.Println()
	fmt.Printf("Is 6 a perfect number? %t", perfectNumber(6))
}
