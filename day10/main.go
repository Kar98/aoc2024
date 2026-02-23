package day10

import (
	"aoc/common"
	"aoc/day10/hiker"
	_ "embed"
)

//go:embed input.txt
var input string

func Solve(part int) int {
	score := 0
	if part == 1 {
		total := []hiker.Coordinate{}
		hike2 := hiker.NewHiker(input)
		points := hike2.GetStartingPoints()

		for _, point := range points {
			endPoints, _ := hike2.Hike(point, []hiker.Coordinate{})
			total = common.GetUniques(endPoints)
			score += len(total)
		}
	}

	return score
}
