package day4

import (
	"aoc/day4/scanner"
	_ "embed"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func Main(puzzle int) int {
	if puzzle == 1 {
		inputData := scanner.FileToInput(input)
		scans := scanner.NewScanner(inputData)
		return scans.ScanForXmas(0, len(inputData)-1)
	}
	return 0
}
