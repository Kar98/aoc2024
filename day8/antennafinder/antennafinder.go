package antennafinder

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

type AntennaFinder struct {
	cityMap     [][]string
	antinodeMap [][]string
}

type Coordinate struct {
	r int // row
	c int // column
}

var ErrOutofBounds = errors.New("out of bounds")

func FileToInput(file string) [][]string {
	lines := strings.Split(file, "\n")
	output := make([][]string, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, "")
		output[i] = nums
	}

	return output
}

func NewFinder(file string) AntennaFinder {
	cityMap := FileToInput(file)
	antiMap := make([][]string, len(cityMap))
	for i := range antiMap {
		antiMap[i] = make([]string, len(cityMap[0]))
		copy(antiMap[i], cityMap[i])
	}
	return AntennaFinder{
		cityMap:     cityMap,
		antinodeMap: antiMap,
	}
}

func (f *AntennaFinder) GetAntennas() []string {
	uniques := []string{}
	for r := range f.cityMap {
		for c := range f.cityMap[r] {
			char := f.cityMap[r][c]
			if char == "." {
				continue
			}
			if !slices.Contains(uniques, char) {
				uniques = append(uniques, char)
			}
		}
	}

	return uniques
}

func (f *AntennaFinder) GetCoordinates(charToFind string) []Coordinate {
	coords := []Coordinate{}
	for r := range f.cityMap {
		for c := range f.cityMap[r] {
			char := f.cityMap[r][c]
			if char == charToFind {
				coords = append(coords, Coordinate{r: r, c: c})
			}
		}
	}
	return coords
}

func (f *AntennaFinder) GetAntinodes(coords []Coordinate) {
	for baseI := 0; baseI < len(coords); baseI++ {
		baseCoord := coords[baseI]

		for _, coord := range coords {
			if baseCoord.r == coord.r && baseCoord.c == coord.c {
				// Found same point, skip
				continue
			}
			// place below
			rdiff := coord.r - baseCoord.r // 2 - 1 =  1
			cdiff := coord.c - baseCoord.c // 4 - 7 = -3
			f.placeAntinode(coord.r+rdiff, coord.c+cdiff)
			// place above
			f.placeAntinode(baseCoord.r+(-1*rdiff), baseCoord.c+(-1*cdiff))
		}
	}
}

func (f *AntennaFinder) GetAntinodesByGrid(coords []Coordinate) {
	for baseI := 0; baseI < len(coords); baseI++ {
		baseCoord := coords[baseI]

		for _, coord := range coords {
			if baseCoord.r == coord.r && baseCoord.c == coord.c {
				// Found same point, count yourself and end
				f.placeAntinode(baseCoord.r, baseCoord.c)
				continue
			}
			// place below
			rdiff := coord.r - baseCoord.r // 2 - 1 =  1
			cdiff := coord.c - baseCoord.c // 4 - 7 = -3
			currentRDiff := 0
			currentCDiff := 0
			var err error
			for err != ErrOutofBounds {
				currentRDiff += rdiff
				currentCDiff += cdiff
				err = f.placeAntinode(coord.r+currentRDiff, coord.c+currentCDiff)
			}
			// place above
			currentRDiff = 0
			currentCDiff = 0
			for err != ErrOutofBounds {
				currentRDiff += rdiff
				currentCDiff += cdiff
				err = f.placeAntinode(baseCoord.r+(-1*currentRDiff), baseCoord.c+(-1*currentCDiff))
			}
		}
	}
}

func (f *AntennaFinder) CountAntinodes() int {
	count := 0
	for r := range f.antinodeMap {
		for c := range f.antinodeMap[r] {
			if f.antinodeMap[r][c] == "#" {
				count++
			}
		}
	}
	return count
}

func (f *AntennaFinder) placeAntinode(row, col int) error {
	if row < 0 || row > len(f.antinodeMap)-1 {
		return ErrOutofBounds
	}
	if col < 0 || col > len(f.antinodeMap[0])-1 {
		return ErrOutofBounds
	}
	f.antinodeMap[row][col] = "#"
	return nil
}

func (f *AntennaFinder) PrintAntinode() {
	for _, row := range f.antinodeMap {
		fmt.Println(row)
	}
}
