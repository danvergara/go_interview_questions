package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	unsorted := []int{7, 1, 8, 2, 9, 4, 6, 5}

	want := []int{1, 2, 4, 5, 6, 7, 8, 9}
	got := Sort(unsorted)

	assert.Equal(t, got, want)
}
