package main

import "time"

func ageCalculator(y, m, d int) int {
	now := time.Now()
	dob := time.Date(y, time.Month(m), d, 0, 0, 0, 0, now.Location())

	age := now.Sub(dob)

	return (int(age.Hours()) / 24) / 365
}
