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
			if len(coords) == 1 {
				continue // skip if only an antenna only has the 1 location
			}
			finder.GetAntinodes(coords)
		}
		return finder.CountAntinodes()
	}
	finder := antennafinder.NewFinder(input)
	antennas := finder.GetAntennas()
	for _, antenna := range antennas {
		coords := finder.GetCoordinates(antenna)
		if len(coords) == 1 {
			continue // skip if only an antenna only has the 1 location
		}
		finder.GetAntinodesByGrid(coords)
	}
	return finder.CountAntinodes()
}
