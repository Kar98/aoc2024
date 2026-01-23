package day8

import (
	"aoc/day8/antennafinder"
	_ "embed"
)

//go:embed input.txt
var input string

func Main(puzzle int) int {
	if puzzle == 1 {
		finder := antennafinder.NewFinder(input)
		antennas := finder.GetAntennas()
		for _, antenna := range antennas {
			coords := finder.GetCoordinates(antenna)
			finder.GetAntinodes(coords)
		}
		return finder.CountAntinodes()
	}
	return 0
}
