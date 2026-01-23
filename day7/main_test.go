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
	operatorMatrix = make(map[int][]string)
	result, err := FileToInput(example)
	assert.NoError(t, err)

	base := 2
	assert.True(t, isValidOperation(result[0], base))
	assert.True(t, isValidOperation(result[1], base))
	assert.True(t, isValidOperation(result[8], base))

	assert.False(t, isValidOperation(result[2], base))
	assert.False(t, isValidOperation(result[3], base))
	assert.False(t, isValidOperation(result[4], base))

	// Check V2 conditions
	operatorMatrix = make(map[int][]string)
	base = 3
	assert.True(t, isValidOperation(result[3], base))
	assert.True(t, isValidOperation(result[3], base))
	assert.True(t, isValidOperation(result[4], base))
}

func TestGenerateOperators(t *testing.T) {
	assert.Equal(t, []string{"+", "x"}, generateOperators([]int64{10, 19}, 2))
	assert.Equal(t, []string{"++", "+x", "x+", "xx"}, generateOperators([]int64{81, 40, 27}, 2))
	assert.Equal(t, []string{"+++", "++x", "+x+", "+xx", "x++", "x+x", "xx+", "xxx"}, generateOperators([]int64{81, 40, 27, 10}, 2))
	operators := generateOperators([]int64{100058717, 58, 44, 628, 7, 145, 31, 9, 8}, 2)
	assert.Contains(t, operators, "++xxxxx+")
	assert.Contains(t, operators, "++++++++")
}

func TestExample(t *testing.T) {
	rows, err := FileToInput(example)
	assert.NoError(t, err)

	var total int64
	for _, row := range rows {
		if isValidOperation(row, 2) {
			total += row[0]
		}
	}

	assert.Equal(t, int64(3749), total)
}

func TestExampleV2(t *testing.T) {
	rows, err := FileToInput(example)
	assert.NoError(t, err)

	var total int64
	for _, row := range rows {
		if isValidOperation(row, 3) {
			total += row[0]
		}
	}

	assert.Equal(t, int64(11387), total)
}

func TestMergeNumbers(t *testing.T) {
	output, err := mergeNumbers(100, 10)
	assert.NoError(t, err)
	assert.Equal(t, int64(10010), output)
}
