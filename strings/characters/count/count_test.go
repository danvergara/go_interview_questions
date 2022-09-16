package count

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	want := map[string]int{
		"c": 2,
		"l": 1,
		"o": 1,
		"d": 1,
		"i": 1,
		"n": 1,
		"g": 1,
	}

	got, _ := Count("clcoding")

	assert.Equal(t, got, want)
}
