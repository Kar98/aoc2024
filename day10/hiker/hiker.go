package hiker

import (
	"errors"
	"strconv"
	"strings"
)

type Hiker struct {
	hikingMap       [][]string
	currentPosition Coordinate
	cBbox           int // column bounding box
	rBbox           int // row bounding box
}

type UniqueHiker struct {
	Hiker
}

type Coordinate struct {
	R int // row
	C int // column
}

var (
	ErrInvalid = errors.New("invalid walk location")
	ErrEnd     = errors.New("end")
)

func FileToInput(file string) [][]string {
	lines := strings.Split(file, "\n")
	output := make([][]string, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, "")
		output[i] = nums
	}

	return output
}

func NewHiker(file string) Hiker {
	input := FileToInput(file)
	return Hiker{
		hikingMap: input,
		cBbox:     len(input[0]) - 1,
		rBbox:     len(input) - 1,
	}
}

func NewUniqueHiker(file string) UniqueHiker {
	input := FileToInput(file)
	return UniqueHiker{Hiker{
		hikingMap: input,
		cBbox:     len(input[0]) - 1,
		rBbox:     len(input) - 1,
	}}
}

func (h *Hiker) GetStartingPoints() []Coordinate {
	positions := []Coordinate{}
	for ri, r := range h.hikingMap {
		for ci := range r {
			if h.hikingMap[ri][ci] == "0" {
				positions = append(positions, Coordinate{R: ri, C: ci})
			}
		}
	}
	return positions
}

func (h *UniqueHiker) Hike(pos Coordinate, resetPos Coordinate, totalFound int) (int, error) {
	// From point r, c
	i, err := strconv.ParseInt(h.hikingMap[pos.R][pos.C], 10, 32)
	if err != nil {
		panic("could not parse: " + h.hikingMap[pos.R][pos.C])
	}
	currentHeight := int(i)
	if currentHeight == 9 {
		return totalFound + 1, ErrEnd
	}
	// Look at each cardinal direction, find number of valid directions
	canGoNorth := h.isValidDirection(pos, Coordinate{R: pos.R - 1, C: pos.C})
	canGoSouth := h.isValidDirection(pos, Coordinate{R: pos.R + 1, C: pos.C})
	canGoWest := h.isValidDirection(pos, Coordinate{R: pos.R, C: pos.C - 1})
	canGoEast := h.isValidDirection(pos, Coordinate{R: pos.R, C: pos.C + 1})
	if canGoNorth {
		//fmt.Println("can go north")
		totalFound, err = h.Hike(Coordinate{R: pos.R - 1, C: pos.C}, pos, totalFound)
	}
	if canGoSouth {
		//fmt.Println("can go south")
		totalFound, err = h.Hike(Coordinate{R: pos.R + 1, C: pos.C}, pos, totalFound)
	}
	if canGoWest {
		//fmt.Println("can go west")
		totalFound, err = h.Hike(Coordinate{R: pos.R, C: pos.C - 1}, pos, totalFound)
	}
	if canGoEast {
		//fmt.Println("can go east")
		totalFound, err = h.Hike(Coordinate{R: pos.R, C: pos.C + 1}, pos, totalFound)
	}

	return totalFound, nil
}

func (h *Hiker) Hike(pos Coordinate, foundHikes []Coordinate) ([]Coordinate, error) {
	var currentHeight int
	//fmt.Printf("r%d c%d\n", position.R, position.C)

	// From point r, c
	i, err := strconv.ParseInt(h.hikingMap[pos.R][pos.C], 10, 32)
	if err != nil {
		return nil, err
	}
	currentHeight = int(i)
	if currentHeight == 9 {
		foundHikes = append(foundHikes, Coordinate{R: pos.R, C: pos.C})
		return foundHikes, ErrEnd
	}

	// Look at each cardinal direction, find number of valid directions
	canGoNorth := h.isValidDirection(pos, Coordinate{R: pos.R - 1, C: pos.C})
	canGoSouth := h.isValidDirection(pos, Coordinate{R: pos.R + 1, C: pos.C})
	canGoWest := h.isValidDirection(pos, Coordinate{R: pos.R, C: pos.C - 1})
	canGoEast := h.isValidDirection(pos, Coordinate{R: pos.R, C: pos.C + 1})
	if canGoNorth {
		//fmt.Println("can go north")
		foundHikes, err = h.Hike(Coordinate{R: pos.R - 1, C: pos.C}, foundHikes)
	}
	if canGoSouth {
		//fmt.Println("can go south")
		foundHikes, err = h.Hike(Coordinate{R: pos.R + 1, C: pos.C}, foundHikes)
	}
	if canGoWest {
		//fmt.Println("can go west")
		foundHikes, err = h.Hike(Coordinate{R: pos.R, C: pos.C - 1}, foundHikes)
	}
	if canGoEast {
		//fmt.Println("can go east")
		foundHikes, err = h.Hike(Coordinate{R: pos.R, C: pos.C + 1}, foundHikes)
	}

	return foundHikes, nil
}

func (h *Hiker) isValidDirection(currentPosition, direction Coordinate) bool {
	if direction.C < 0 || direction.R < 0 {
		return false
	}
	if direction.C > h.cBbox {
		return false
	}
	if direction.R > h.rBbox {
		return false
	}
	// Found in examples only, real input will always have numbers
	if h.hikingMap[direction.R][direction.C] == "." {
		return false
	}
	diff := h.toInt(direction) - h.toInt(currentPosition)
	if diff != 1 {
		return false
	}
	return true
}

func (h *Hiker) toInt(direction Coordinate) int {
	i, err := strconv.ParseInt(h.hikingMap[direction.R][direction.C], 10, 32)
	if err != nil {
		panic("non number found in hikingMap")
	}
	return int(i)
}
