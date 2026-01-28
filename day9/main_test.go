package day9

import (
	"aoc/day9/diskbuilder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	builder := diskbuilder.NewSimpleDisk("12345")
	err := builder.BuildDisk()
	assert.NoError(t, err)
	out := builder.PrintDisk()
	assert.Equal(t, "0..111....22222", out)

	builder2 := diskbuilder.NewSimpleDisk("2333133121414131402")
	err = builder2.BuildDisk()
	assert.NoError(t, err)
	out = builder2.PrintDisk()
	assert.Equal(t, "00...111...2...333.44.5555.6666.777.888899", out)
}

func TestSorting(t *testing.T) {
	builder := diskbuilder.NewSimpleDisk("2333133121414131402")
	err := builder.BuildDisk()
	assert.NoError(t, err)
	builder.Sort()
	out := builder.PrintDisk()
	assert.Equal(t, "0099811188827773336446555566..............", out)
	checksum := builder.CalculateChecksum()
	assert.Equal(t, int64(1928), checksum)
}
