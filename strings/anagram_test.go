package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnagram(t *testing.T) {
	assert.True(t, anagram("cinema", "iceman"))
	assert.True(t, anagram("cool", "loco"))
	assert.False(t, anagram("men", "women"))
	assert.True(t, anagram("python", "pythno"))
}
