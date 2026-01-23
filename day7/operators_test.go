package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOperatorStr(t *testing.T) {
	slc := createOperatorStr(10, 3, 3)
	assert.Equal(t, "x+x", slc)
}
