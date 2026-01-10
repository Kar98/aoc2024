package day5

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileToInput(t *testing.T) {
	output := FileToInput(example)
	assert.Equal(t, 6, len(output.rules))
	assert.Len(t, output.rules["29"], 1)
	assert.Equal(t, "75,47,61,53,29", output.pages[0])
}

func TestPrintRules(t *testing.T) {
	output := FileToInput(example)
	for x, y := range output.rules {
		fmt.Printf("x: %s . y: %s\n", x, strings.Join(y, ","))
	}
}

func TestExample(t *testing.T) {
	res := Example(1)
	assert.Equal(t, 143, res)

	res = Example(2)
	assert.Equal(t, 123, res)
}

func TestRule(t *testing.T) {
	puzzle := FileToInput(example)
	got, _, err := IsValidPage("75,97,47,61,53", puzzle.rules)
	assert.NoError(t, err)
	assert.False(t, got)

	rules1 := map[string][]string{}
	rules1["10"] = []string{"11"}
	got, _, err = IsValidPage("11,10", rules1)
	assert.NoError(t, err)
	assert.False(t, got)

	rules2 := map[string][]string{}
	rules2["15"] = []string{"14"}
	got, _, err = IsValidPage("10,11,12,13,14,15", rules2)
	assert.False(t, got)
	assert.NoError(t, err)
	assert.False(t, got)
}

func TestPageSorting(t *testing.T) {
	/*
		75,97,47,61,53 becomes 97,75,47,61,53
		61,13,29 becomes 61,29,13
		97,13,75,29,47 becomes 97,75,47,29,13
	*/
	puzzle := FileToInput(example)
	pageNums := strings.Split("75,97,47,61,53", ",")
	SortPage(pageNums, puzzle.rules)
	assert.Equal(t, strings.Split("97,75,47,61,53", ","), pageNums)

	pageNums = strings.Split("61,13,29", ",")
	SortPage(pageNums, puzzle.rules)
	assert.Equal(t, strings.Split("61,29,13", ","), pageNums)

	pageNums = strings.Split("97,13,75,29,47", ",")
	SortPage(pageNums, puzzle.rules)
	assert.Equal(t, strings.Split("97,75,47,29,13", ","), pageNums)
}

func TestMoveInSlice(t *testing.T) {
	pageNums := strings.Split("97,13,29,75,47", ",")
	modified, err := moveInSlice(pageNums, 3, 2)
	assert.NoError(t, err)
	assert.Equal(t, strings.Split("97,13,75,29,47", ","), modified)
	assert.Equal(t, strings.Split("97,13,29,75,47", ","), pageNums) // original is unchanged

	modified, err = moveInSlice(pageNums, 3, 1)
	assert.Equal(t, strings.Split("97,75,13,29,47", ","), modified)

	modified, err = moveInSlice(pageNums, 3, 0)
	assert.Equal(t, strings.Split("75,97,13,29,47", ","), modified)
}

func TestEditSlice(t *testing.T) {
	page1 := "97,13,29,75,47"
	pageNums := strings.Split(page1, ",")
	modified, err := editSlice(pageNums, 3, 2)
	assert.NoError(t, err)
	assert.Equal(t, strings.Split("97,13,75,29,47", ","), modified)
	assert.NotEqual(t, strings.Split(page1, ","), modified)
}
