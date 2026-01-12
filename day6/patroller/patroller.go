package patroller

import (
	"errors"
	"strings"
)

type Direction string

const (
	Up    Direction = "u"
	Down  Direction = "d"
	Left  Direction = "l"
	Right Direction = "r"
)

var (
	ErrObstacle = errors.New("obstacle")
	ErrEnd      = errors.New("end")
)

type Patroller struct {
	labArea   [][]string
	direction Direction
	r         int // row
	c         int // coloumn
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
	r, c := startPoint(puzzleInput)
	return Patroller{
		labArea:   puzzleInput,
		direction: Up,
		r:         r,
		c:         c,
	}
}

func (p *Patroller) CurrentPosition() (int, int) {
	return p.r, p.c
}

func (p *Patroller) WalkUp() error {
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

func (p *Patroller) WalkDown() error {
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
func (p *Patroller) WalkRight() error {
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
func (p *Patroller) WalkLeft() error {
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
		return p.WalkUp
	}
	if p.direction == Down {
		return p.WalkDown
	}
	if p.direction == Right {
		return p.WalkRight
	}
	if p.direction == Left {
		return p.WalkLeft
	}
	return p.WalkUp
}

func (p *Patroller) changeDirection() (Direction, error) {
	if p.direction == Up {
		p.walk = p.WalkRight
		return Right, nil
	}
	if p.direction == Right {
		p.walk = p.WalkDown
		return Down, nil
	}
	if p.direction == Down {
		p.walk = p.WalkLeft
		return Left, nil
	}
	if p.direction == Left {
		p.walk = p.WalkUp
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
		if failSafe > 100000 {
			return 0, errors.New("exceeded failsafe threshold")
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
