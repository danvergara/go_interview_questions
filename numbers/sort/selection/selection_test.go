package selection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	unsorted := []int{4, 6, 9, 8, 1, 7, 3}

	want := []int{1, 3, 4, 6, 7, 8, 9}
	got := Sort(unsorted)

	assert.Equal(t, got, want)
}
