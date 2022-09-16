package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAPerfectNumber(t *testing.T) {
	assert.True(t, IsAPerfectNumber(6))
}
