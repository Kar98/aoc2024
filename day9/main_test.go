package day9

import (
	"aoc/day9/diskbuilder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStart(t *testing.T) {
	/// 12345
	// file, space, file, sapce, file

	// 0..111....22222

	builder := diskbuilder.NewDiskBuilder("12345")
	err := builder.Build()
	assert.NoError(t, err)
	builder.PrintDisk()

}
