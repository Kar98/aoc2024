package day5

import (
	_ "embed"
	"errors"
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
	puzzleData := FileToInput(example)
	if part == 1 {
		total, _ := solvePuzzlePart1(puzzleData)
		return total
	} else {
		return solvePuzzlePart2(puzzleData)
	}
}

func Main(part int) int {
	puzzleData := FileToInput(input)
	if part == 1 {
		total, _ := solvePuzzlePart1(puzzleData)
		return total
	} else {
		return solvePuzzlePart2(puzzleData)
	}
}

func solvePuzzlePart1(puzzle Puzzle) (int, []string) {
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

func solvePuzzlePart2(puzzle Puzzle) int {
	total := 0
	pagesToSort := []string{}
	for _, page := range puzzle.pages {
		isValid, _, err := IsValidPage(page, puzzle.rules)
		if err != nil {
			fmt.Println(err.Error())
			return total
		}
		// Only count the incorrect added pages
		if !isValid {
			pagesToSort = append(pagesToSort, page)
		}
	}

	for _, page := range pagesToSort {
		pageNums := strings.Split(page, ",")
		SortPage(pageNums, puzzle.rules)
		middleIdx := len(pageNums) / 2
		parsedInt, err := strconv.ParseInt(pageNums[middleIdx], 10, 32)
		if err != nil {
			fmt.Println(err.Error())
			return total
		}
		total += int(parsedInt)
	}

	return total
}

func moveInSlice(s []string, fromIdx int, toIdx int) ([]string, error) {
	if fromIdx <= toIdx {
		return []string{}, errors.New("attempting to copy forwards")
	}
	newSlice := make([]string, len(s))
	copy(newSlice, s)
	tmpVal := s[fromIdx]
	if toIdx == 0 {
		newSlice = slices.Delete(newSlice, fromIdx, fromIdx+1)
		return slices.Insert(newSlice, 0, tmpVal), nil
	}
	for i := fromIdx; i >= toIdx; i-- {
		newSlice[i] = newSlice[i-1]
	}
	newSlice[toIdx] = tmpVal
	return newSlice, nil
}

func editSlice(s []string, fromIdx int, toIdx int) ([]string, error) {
	if fromIdx <= toIdx {
		return []string{}, errors.New("attempting to copy forwards")
	}
	tmpVal := s[fromIdx]
	if toIdx == 0 {
		s = slices.Delete(s, fromIdx, fromIdx+1)
		return slices.Insert(s, 0, tmpVal), nil
	}
	for i := fromIdx; i >= toIdx; i-- {
		s[i] = s[i-1]
	}
	s[toIdx] = tmpVal
	return s, nil
}

func SortPage(pageNums []string, rules map[string][]string) []string {
	for i := 0; i < len(pageNums); i++ {
		pageNum := pageNums[i]
		rule, ok := rules[pageNum]
		// If no rule exists, then we don't need to check
		if !ok {
			continue
		}
		// Go through every number in the rule, and ensure that all previous indexs do not match a rule number.
		// If a rule number is matched, then place the current index before the matched number.
		previousIndexes := pageNums[0:i]
		for _, ruleNum := range rule {
			found, at := ruleInPreviousNums(previousIndexes, ruleNum)
			if found {
				editSlice(pageNums, i, at)
				return SortPage(pageNums, rules)
			}
		}
	}

	return pageNums
}

func ruleInPreviousNums(previousNums []string, ruleNum string) (bool, int) {
	for i, prevNum := range previousNums {
		if prevNum == ruleNum {
			return true, i
		}
	}
	return false, -1
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
