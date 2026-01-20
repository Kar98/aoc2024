package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSlice(t *testing.T) {
	slc := createSlice(10, 3)
	assert.Equal(t, []string{"x", "+", "x"}, slc)
}
