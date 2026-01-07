package day5

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

type Puzzle struct {
	rules map[string][]string
	pages []string
}

func FileToInput(input string) Puzzle {
	ruleMapping := map[string][]string{}
	parts := strings.Split(input, "\n\n")
	// Rules
	rules := strings.SplitSeq(parts[0], "\n")
	for rule := range rules {
		splits := strings.Split(rule, "|")
		xNumber := splits[0]
		yNumber := splits[1]
		_, ok := ruleMapping[xNumber]
		if !ok {
			ruleMapping[xNumber] = []string{}
		}
		ruleMapping[xNumber] = append(ruleMapping[xNumber], yNumber)
	}
	// Pages
	pages := strings.Split(parts[1], "\n")
	return Puzzle{
		rules: ruleMapping,
		pages: pages,
	}
}

func Example(part int) int {
	if part == 1 {
		puzzleData := FileToInput(example)
		total, _ := solvePuzzle(puzzleData)
		return total
	} else {
		return 0
	}
}

func Main(part int) int {

	// 47 must be before all the numbers on the right
	// all numbers in the rule, must not be found before the RuleNumber
	// If pages meet the conditions, then grab the middle number and add it to the tally

	if part == 1 {
		puzzleData := FileToInput(input)
		total, _ := solvePuzzle(puzzleData)
		return total
	} else {
		return 0
	}
}

func solvePuzzle(puzzle Puzzle) (int, []string) {
	total := 0
	validPages := []string{}
	for _, page := range puzzle.pages {
		ok, iValue, err := IsValidPage(page, puzzle.rules)
		if err != nil {
			fmt.Println(err.Error())
			return total, validPages
		}
		if ok {
			validPages = append(validPages, page)
			total += iValue
		}
	}
	return total, validPages
}

func IsValidPage(page string, rules map[string][]string) (bool, int, error) {
	pageNums := strings.Split(page, ",")
	isValid := true
	middleIdx := len(pageNums) / 2
	// Never need to check the first instance
	for i := 1; i < len(pageNums); i++ {
		pageNum := pageNums[i]
		rule, ok := rules[pageNum]
		// If no rule exists, then we don't need to check
		if !ok {
			continue
		}
		if !checkValidPage(rule, i, &pageNums) {
			isValid = false
		}
	}
	result, err := strconv.ParseInt(pageNums[middleIdx], 10, 32)
	return isValid, int(result), err
}

func checkValidPage(rule []string, index int, pageNums *[]string) bool {
	for i := 0; i < index; i++ {
		if slices.Contains(rule, (*pageNums)[i]) {
			return false
		}
	}
	return true
}
