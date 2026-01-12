package day6

import (
	"aoc/day6/patroller"
	_ "embed"
	"fmt"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func Main(part int) int {
	file := patroller.LoadFile(input)
	if part == 1 {
		patrol := patroller.NewPatroller(file)
		total, err := patrol.GoWalking()
		if err != nil {
			fmt.Println(err.Error())
		}
		return total
	}

	return 0
}
