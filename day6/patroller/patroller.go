package patroller

import (
	"errors"
	"strings"
)

type Direction string

// RowColumn
type RC struct {
	R int
	C int
}

const (
	Up    Direction = "u"
	Down  Direction = "d"
	Left  Direction = "l"
	Right Direction = "r"
)

var (
	ErrObstacle = errors.New("obstacle")
	ErrEnd      = errors.New("end")
	ErrLoop     = errors.New("stuck in loop")
)

type Patroller struct {
	labArea   [][]string
	direction Direction
	r         int // row
	startingR int
	c         int // coloumn
	startingC int
	walk      func() error
}

func LoadFile(input string) [][]string {
	var output [][]string
	output = [][]string{}
	rows := strings.SplitSeq(input, "\n")
	for row := range rows {
		splits := strings.Split(row, "")
		output = append(output, splits)
	}
	return output
}

func startPoint(input [][]string) (int, int) {
	for ri := range input {
		for ci := range input[ri] {
			if input[ri][ci] == "^" {
				return ri, ci
			}
		}
	}
	return 0, 0
}

func NewPatroller(puzzleInput [][]string) Patroller {
	var labArea [][]string
	labArea = make([][]string, len(puzzleInput))
	for i := range puzzleInput {
		labArea[i] = make([]string, len(puzzleInput[i]))
		copy(labArea[i], puzzleInput[i])
	}

	r, c := startPoint(labArea)
	return Patroller{
		labArea:   labArea,
		direction: Up,
		r:         r,
		c:         c,
		startingR: r,
		startingC: c,
	}
}

func (p *Patroller) CurrentPosition() (int, int) {
	return p.r, p.c
}

func (p *Patroller) GetLabArea() [][]string {
	return p.labArea
}

func (p *Patroller) walkUp() error {
	if p.r-1 < 0 {
		return ErrEnd
	}
	if p.labArea[p.r-1][p.c] == "#" {
		return ErrObstacle
	}

	p.r--
	p.labArea[p.r][p.c] = "X"

	return nil
}

func (p *Patroller) walkDown() error {
	if p.r+1 >= len(p.labArea) {
		return ErrEnd
	}
	if p.labArea[p.r+1][p.c] == "#" {
		return ErrObstacle
	}

	p.r++
	p.labArea[p.r][p.c] = "X"

	return nil
}
func (p *Patroller) walkRight() error {
	if p.c+1 >= len(p.labArea[0]) {
		return ErrEnd
	}
	if p.labArea[p.r][p.c+1] == "#" {
		return ErrObstacle
	}

	p.c++
	p.labArea[p.r][p.c] = "X"

	return nil
}
func (p *Patroller) walkLeft() error {
	if p.c-1 < 0 {
		return ErrEnd
	}
	if p.labArea[p.r][p.c-1] == "#" {
		return ErrObstacle
	}

	p.c--
	p.labArea[p.r][p.c] = "X"

	return nil
}

func (p *Patroller) getWalkFunc() func() error {
	if p.direction == Up {
		return p.walkUp
	}
	if p.direction == Down {
		return p.walkDown
	}
	if p.direction == Right {
		return p.walkRight
	}
	if p.direction == Left {
		return p.walkLeft
	}
	return p.walkUp
}

func (p *Patroller) changeDirection() (Direction, error) {
	if p.direction == Up {
		p.walk = p.walkRight
		return Right, nil
	}
	if p.direction == Right {
		p.walk = p.walkDown
		return Down, nil
	}
	if p.direction == Down {
		p.walk = p.walkLeft
		return Left, nil
	}
	if p.direction == Left {
		p.walk = p.walkUp
		return Up, nil
	}
	return "", errors.New("unknown direction change")
}

func (p *Patroller) countVisitedTiles() int {
	total := 0
	for ri := range p.labArea {
		for ci := range p.labArea[ri] {
			if p.labArea[ri][ci] == "X" {
				total++
			}
		}
	}
	return total
}

func (p *Patroller) GetXPositions() []RC {
	output := []RC{}
	for r := range p.labArea {
		for c := range p.labArea[r] {
			// Dont count the starting position as we don't want to put an obstacle there.
			if p.startingC == c && p.startingR == r {
				continue
			}
			if p.labArea[r][c] == "X" {
				output = append(output, RC{R: r, C: c})
			}
		}
	}
	return output
}

func (p *Patroller) AddObstacle(rc RC) {
	p.labArea[rc.R][rc.C] = "#"
}

func (p *Patroller) GoWalking() (int, error) {
	failSafe := 0
	// Get current direction
	// After every walk, the current tile is marked with X and continue
	// Keep walking until # is hit
	// When # is hit, then turn right, and keep walking
	// Keep walking until end
	p.walk = p.getWalkFunc()
	// Mark starting position as visited
	p.labArea[p.r][p.c] = "X"

	var walkErr error
	for walkErr != ErrEnd {
		if failSafe > 10000 {
			return 0, ErrLoop
		}
		failSafe++

		walkErr = p.walk()
		if walkErr == ErrObstacle {
			newDir, err := p.changeDirection()
			p.direction = newDir
			if err != nil {
				return 0, err
			}
		}
	}

	return p.countVisitedTiles(), nil
}
