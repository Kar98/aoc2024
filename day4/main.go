package day4

import (
	"aoc/day4/scanner"
	_ "embed"
	"sync"
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

// Quick and dirty hardcode to see how much faster with goroutines
func MainWithGoroutines(puzzle int) int {
	total := 0
	inputData := scanner.FileToInput(input)
	scans := scanner.NewScanner(inputData)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		total += scans.ScanForXmas(0, 49)
	}()
	go func() {
		defer wg.Done()
		total += scans.ScanForXmas(50, 100)
	}()
	go func() {
		defer wg.Done()
		total += scans.ScanForXmas(101, 139)
	}()

	wg.Wait()
	return total
}
