package main

func perfectNumber(number int) bool {
	sum := 0
	for x := 1; x < number; x++ {
		if number%x == 0 {
			sum += x
		}
	}

	return sum == number
}
