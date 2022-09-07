package main

import (
	"encoding/json"
	"fmt"
)

func countCharacters(s string) {
	count := make(map[string]int)

	for _, i := range s {
		count[string(i)]++
	}

	b, err := json.MarshalIndent(count, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%s\n", string(b))
}
