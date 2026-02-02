package diskbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplexDiskBuilder(t *testing.T) {
	builder := NewComplexDisk("12345")
	err := builder.BuildDisk()
	assert.NoError(t, err)
	assert.Equal(t, 10, builder.lookup[2].startPos)
}
