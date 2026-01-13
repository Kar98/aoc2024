package day6

import (
	"aoc/day6/patroller"
	"errors"
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

func TestExamplePart2(t *testing.T) {
	file := patroller.LoadFile(example)
	patrol := patroller.NewPatroller(file)
	total, err := patrol.GoWalking()
	assert.NoError(t, err)
	assert.Equal(t, 41, total)
	xs := patrol.GetXPositions()

	loops := 0
	for _, pos := range xs {
		patrol2 := patroller.NewPatroller(file)
		patrol2.AddObstacle(pos)
		_, err := patrol2.GoWalking()
		if errors.Is(err, patroller.ErrLoop) {
			loops++
		}
	}
	assert.Equal(t, 6, loops)
}

func TestOriginalSliceNotEdited(t *testing.T) {
	file := patroller.LoadFile(example)
	patrol := patroller.NewPatroller(file)
	initialR, initialC := patrol.CurrentPosition()
	_, err := patrol.GoWalking()
	assert.NoError(t, err)
	newLab := patrol.GetLabArea()

	assert.Equal(t, "^", file[initialR][initialC])
	assert.Equal(t, "X", newLab[initialR][initialC])
}

func TestExampleXPositions(t *testing.T) {
	file := patroller.LoadFile(example)
	patrol := patroller.NewPatroller(file)
	_, err := patrol.GoWalking()
	assert.NoError(t, err)
	assert.Len(t, patrol.GetXPositions(), 40) // X minus starting pos
}

func TestExampleWithObstacle(t *testing.T) {
	file := patroller.LoadFile(example)
	// Put obstacle in way and cause loop
	patrol2 := patroller.NewPatroller(file)
	patrol2.AddObstacle(patroller.RC{R: 6, C: 3})
	_, err := patrol2.GoWalking()
	assert.ErrorIs(t, patroller.ErrLoop, err)
}
