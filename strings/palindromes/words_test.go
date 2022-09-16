package words

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeWords(t *testing.T) {
	palindromeCandidates := "Civic level noon support career"

	want := []string{"civic", "level", "noon"}
	got := Palindromes(palindromeCandidates)

	assert.EqualValues(t, want, got)
}
