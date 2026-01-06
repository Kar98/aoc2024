package scanner

import (
	"errors"
	"fmt"
	"strings"
)

type Scanner struct {
	bbox  BoundingBox
	input [][]string
}

func FileToInput(input string) [][]string {
	output := [][]string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		output = append(output, strings.Split(line, ""))
	}
	return output
}

func Main(_ int) int {
	input := [][]string{{"1", "2", "3"}, {"4", "5", "6"}}
	fmt.Println(input[0])
	fmt.Println(input[1])
	fmt.Println(input[1][2])
	return 0
}

func Example() {

}

func NewScanner(input [][]string) Scanner {
	return Scanner{
		bbox: BoundingBox{
			lBox: 0,
			rBox: len(input[0]) - 1,
			tBox: 0,
			bBox: len(input) - 1,
		},
		input: input,
	}
}

type BoundingBox struct {
	lBox int // Left
	rBox int // Right
	tBox int // Top
	bBox int // Bottom
}

func (s *Scanner) Scan(startingRow int) int {
	xmasCount := 0

	// Cursor scans through the input and will check each orthoganal direction
	for r := startingRow; r < len(s.input); r++ {
		for c := 0; c < len(s.input[r]); c++ {
			// Check top
			out, err := s.Up(r, c)
			if out == "XMAS" && err == nil {
				xmasCount++
			}
			// Check right
			out, err = s.Right(r, c)
			if out == "XMAS" && err == nil {
				xmasCount++
			}
			// Check down
			out, err = s.Down(r, c)
			if out == "XMAS" && err == nil {
				xmasCount++
			}
			// Check left
			out, err = s.Left(r, c)
			if out == "XMAS" && err == nil {
				xmasCount++
			}
			// Check diagonal directions
			out, err = s.UpLeft(r, c)
			if out == "XMAS" && err == nil {
				xmasCount++
			}
			out, err = s.UpRight(r, c)
			if out == "XMAS" && err == nil {
				xmasCount++
			}
			out, err = s.DownLeft(r, c)
			if out == "XMAS" && err == nil {
				xmasCount++
			}
			out, err = s.DownRight(r, c)
			if out == "XMAS" && err == nil {
				xmasCount++
			}
		}
	}
	return xmasCount
}

func (s *Scanner) Up(r int, c int) (string, error) {
	if r-3 < s.bbox.tBox {
		return "", errors.New("out of bounds")
	}
	str := s.input[r][c] + s.input[r-1][c] + s.input[r-2][c] + s.input[r-3][c]
	return str, nil
}

func (s *Scanner) UpLeft(r int, c int) (string, error) {
	if r-3 < s.bbox.tBox || c-3 < s.bbox.lBox {
		return "", errors.New("out of bounds")
	}
	str := s.input[r][c] + s.input[r-1][c-1] + s.input[r-2][c-2] + s.input[r-3][c-3]
	return str, nil
}

func (s *Scanner) UpRight(r int, c int) (string, error) {
	if r-3 < s.bbox.tBox || c+3 > s.bbox.rBox {
		return "", errors.New("out of bounds")
	}
	str := s.input[r][c] + s.input[r-1][c+1] + s.input[r-2][c+2] + s.input[r-3][c+3]
	return str, nil
}

func (s *Scanner) Right(r int, c int) (string, error) {
	if c+3 > s.bbox.rBox {
		return "", errors.New("out of bounds")
	}
	str := s.input[r][c] + s.input[r][c+1] + s.input[r][c+2] + s.input[r][c+3]
	return str, nil
}

func (s *Scanner) Down(r int, c int) (string, error) {
	if r+3 > s.bbox.bBox {
		return "", errors.New("out of bounds")
	}
	str := s.input[r][c] + s.input[r+1][c] + s.input[r+2][c] + s.input[r+3][c]
	return str, nil
}

func (s *Scanner) DownLeft(r int, c int) (string, error) {
	if r+3 > s.bbox.bBox || c-3 < s.bbox.lBox {
		return "", errors.New("out of bounds")
	}
	str := s.input[r][c] + s.input[r+1][c-1] + s.input[r+2][c-2] + s.input[r+3][c-3]
	return str, nil
}

func (s *Scanner) DownRight(r int, c int) (string, error) {
	if r+3 > s.bbox.bBox || c+3 > s.bbox.rBox {
		return "", errors.New("out of bounds")
	}
	str := s.input[r][c] + s.input[r+1][c+1] + s.input[r+2][c+2] + s.input[r+3][c+3]
	return str, nil
}

func (s *Scanner) Left(r int, c int) (string, error) {
	if c-3 < s.bbox.lBox {
		return "", errors.New("out of bounds")
	}
	str := s.input[r][c] + s.input[r][c-1] + s.input[r][c-2] + s.input[r][c-3]
	return str, nil
}
