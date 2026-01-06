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
		scans := scanner.NewScanner(scanner.FileToInput(input))
		return scans.Scan(0)
	}
	return 0
}
