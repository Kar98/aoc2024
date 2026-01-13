package day6

import (
	"aoc/day6/patroller"
	_ "embed"
	"errors"
	"fmt"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func Main(part int) int {
	file := patroller.LoadFile(input)
	total := 0
	if part == 1 {
		patrol := patroller.NewPatroller(file)
		total, err := patrol.GoWalking()
		if err != nil {
			fmt.Println(err.Error())
		}
		return total
	}
	// Execute part1
	patrol := patroller.NewPatroller(file)
	_, err := patrol.GoWalking()
	if err != nil {
		fmt.Println(err.Error())
	}
	// Find all positions of Xs
	xPositions := patrol.GetXPositions()
	fmt.Printf("Total X positions: %d\n", len(xPositions))
	// Iterate through , put 1 # in walking path
	for _, pos := range xPositions {
		patrol2 := patroller.NewPatroller(file)
		patrol2.AddObstacle(pos)
		_, err := patrol2.GoWalking()
		if errors.Is(patroller.ErrLoop, err) {
			total++
		}
	}

	return total
}
