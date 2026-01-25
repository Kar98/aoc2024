package day8

import (
	"aoc/day8/antennafinder"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed example.txt
var example string

func TestFileToInput(t *testing.T) {
	res := antennafinder.FileToInput(example)
	assert.Len(t, res, 12)
	assert.Len(t, res[0], 12)
	assert.Equal(t, ".", res[0][0])
	assert.Equal(t, ".", res[11][11])
	assert.Equal(t, "A", res[9][9])

}

func TestAntennaFinder(t *testing.T) {
	afinder := antennafinder.NewFinder(example)
	assert.Len(t, afinder.GetAntennas(), 2)
}

func TestGetCoords(t *testing.T) {
	afinder := antennafinder.NewFinder(example)
	assert.Len(t, afinder.GetCoordinates("0"), 4)
	assert.Len(t, afinder.GetCoordinates("A"), 3)
}

func TestAntinodePlacement(t *testing.T) {
	afinder := antennafinder.NewFinder(example)
	coords1 := afinder.GetCoordinates("0")
	coords2 := afinder.GetCoordinates("A")
	afinder.GetAntinodes(coords1)
	afinder.GetAntinodes(coords2)
	afinder.PrintAntinode()
	assert.Equal(t, 14, afinder.CountAntinodes())
}

func TestAntinodeGridPlacement(t *testing.T) {
	afinder := antennafinder.NewFinder(example)
	coords1 := afinder.GetCoordinates("0")
	coords2 := afinder.GetCoordinates("A")
	afinder.GetAntinodesByGrid(coords1)
	afinder.GetAntinodesByGrid(coords2)
	afinder.PrintAntinode()
	assert.Equal(t, 34, afinder.CountAntinodes())
}
