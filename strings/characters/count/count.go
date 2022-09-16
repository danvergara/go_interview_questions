package count

import (
	"encoding/json"
)

// Count counts the number ocurrencies in a given word.
func Count(s string) (map[string]int, error) {
	count := make(map[string]int)

	for _, i := range s {
		count[string(i)]++
	}

	_, err := json.MarshalIndent(count, "", "  ")
	if err != nil {
		return nil, err
	}

	return count, err
}
