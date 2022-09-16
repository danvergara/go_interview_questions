package age

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	want := 29
	got := Calculate(1992, 11, 16)

	assert.Equal(t, want, got)
}
