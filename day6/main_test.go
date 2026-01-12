package day6

import (
	"aoc/day6/patroller"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	file := patroller.LoadFile(example)
	assert.Len(t, file, 10)
	t.Log(file[0])
}

func TestExampleStartPoint(t *testing.T) {
	file := patroller.LoadFile(example)
	patrol := patroller.NewPatroller(file)
	y, x := patrol.CurrentPosition()
	assert.Equal(t, 6, y)
	assert.Equal(t, 4, x)
	t.Log(x, y)
}

func TestExample(t *testing.T) {
	file := patroller.LoadFile(example)
	patrol := patroller.NewPatroller(file)
	total, err := patrol.GoWalking()
	assert.NoError(t, err)
	assert.Equal(t, 41, total)
}
