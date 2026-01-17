package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileToInput(t *testing.T) {
	result, err := FileToInput("190: 10 19")
	assert.NoError(t, err)
	assert.Equal(t, [][]int64{{190, 10, 19}}, result)

	result2, err := FileToInput(example)
	assert.NoError(t, err)
	assert.Equal(t, []int64{3267, 81, 40, 27}, result2[1])

	result3, err := FileToInput(input)
	assert.NoError(t, err)
	assert.Equal(t, []int64{31084, 8, 67, 8, 735, 38}, result3[0])
	assert.Equal(t, []int64{40541461584, 5, 81, 9, 4, 32, 43, 283, 228}, result3[1])
}

func TestValidOperationsFunc(t *testing.T) {
	result, err := FileToInput(example)
	assert.NoError(t, err)

	assert.True(t, isValidOperation(result[0]))
	assert.True(t, isValidOperation(result[1]))
	assert.True(t, isValidOperation(result[8]))

	assert.False(t, isValidOperation(result[2]))
	assert.False(t, isValidOperation(result[3]))
	assert.False(t, isValidOperation(result[4]))
}

func TestGenerateOperators(t *testing.T) {
	assert.Equal(t, []string{"0", "1"}, generateOperators([]int64{10, 19}))
	assert.Equal(t, []string{"00", "01", "10", "11"}, generateOperators([]int64{81, 40, 27}))
	assert.Equal(t, []string{"000", "001", "010", "011", "100", "101", "110", "111"}, generateOperators([]int64{81, 40, 27, 10}))
}

func TestPad(t *testing.T) {
	assert.Equal(t, "0001", pad("1", 4))
	assert.Equal(t, "0010", pad("0010", 4))
	assert.Equal(t, "000100", pad("0100", 6))
}

func TestConvertBinToOperators(t *testing.T) {
	res := convertBinToOperators([]string{"01001"})
	assert.Equal(t, "x+xx+", res[0])
}

func TestExample(t *testing.T) {
	rows, err := FileToInput(example)
	assert.NoError(t, err)

	var total int64
	for _, row := range rows {
		if isValidOperation(row) {
			total += row[0]
		}
	}

	assert.Equal(t, int64(3749), total)
}
