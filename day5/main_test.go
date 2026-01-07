package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileToInput(t *testing.T) {
	output := FileToInput(example)
	assert.Equal(t, 6, len(output.rules))
	assert.Len(t, output.rules["29"], 1)
	assert.Equal(t, "75,47,61,53,29", output.pages[0])
}

func TestExample(t *testing.T) {
	res := Main(1)
	assert.Equal(t, 143, res)
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
