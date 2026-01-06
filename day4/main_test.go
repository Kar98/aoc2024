package day4

import (
	"aoc/day4/scanner"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputToFile(t *testing.T) {
	out := scanner.FileToInput(example)
	if len(out) != 10 {
		t.Error("not 10 lines")
	}

	assert.Equal(t, "A", out[3][2])
}

func TestUpOoB(t *testing.T) {
	s := scanner.NewScanner(scanner.FileToInput(example))
	// Negative
	_, err := s.Up(2, 1)
	assert.Error(t, err)
	// Positive
	got, err := s.Up(3, 1)
	assert.NoError(t, err)
	assert.Equal(t, "SMSM", got)
}

func TestRightOoB(t *testing.T) {
	s := scanner.NewScanner(scanner.FileToInput(example))
	// Negative
	_, err := s.Right(0, 7)
	assert.Error(t, err)
	// Positive
	got, err := s.Right(0, 6)
	assert.NoError(t, err)
	assert.Equal(t, "MASM", got)

	got, err = s.Right(0, 0)
	assert.Equal(t, "MMMS", got)
}

func TestDownOoB(t *testing.T) {
	s := scanner.NewScanner(scanner.FileToInput(example))
	// Negative
	_, err := s.Down(7, 5)
	assert.Error(t, err)
	// Positive
	got, err := s.Down(6, 5)
	assert.NoError(t, err)
	assert.Equal(t, "AAXX", got)
}

func TestLeftOoB(t *testing.T) {
	s := scanner.NewScanner(scanner.FileToInput(example))
	// Negative
	_, err := s.Left(1, 2)
	assert.Error(t, err)
	// Positive
	got, err := s.Left(1, 3)
	assert.NoError(t, err)
	assert.Equal(t, "MASM", got)
}

func TestDownLeftOoB(t *testing.T) {
	s := scanner.NewScanner(scanner.FileToInput(example))
	// Negative
	_, err := s.DownLeft(3, 0)
	assert.Error(t, err)
	_, err = s.DownLeft(7, 3)
	assert.Error(t, err)
	// Positive
	got, err := s.DownLeft(6, 3)
	assert.NoError(t, err)
	assert.Equal(t, "MXAM", got)
}

func TestExample(t *testing.T) {
	inputData := scanner.FileToInput(example)
	s := scanner.NewScanner(inputData)
	assert.Equal(t, 18, s.ScanForXmas(0, len(inputData)-1))
}

func TestInBbox(t *testing.T) {
	inputData := scanner.FileToInput(example)
	s := scanner.NewScannerV2(inputData)
	assert.True(t, s.InBbox(1, 1))
	assert.False(t, s.InBbox(1, 0))
	assert.False(t, s.InBbox(0, 1))
	assert.False(t, s.InBbox(0, 0))
	assert.True(t, s.InBbox(8, 8))
	assert.False(t, s.InBbox(8, 9))
	assert.False(t, s.InBbox(9, 8))
	assert.False(t, s.InBbox(9, 9))
	assert.False(t, s.InBbox(10, 10))
}

func TestBackslash(t *testing.T) {
	inputData := scanner.FileToInput(example)
	s := scanner.NewScannerV2(inputData)
	assert.False(t, s.CheckBackSlash(0, 0))
	assert.False(t, s.CheckBackSlash(1, 1))
	assert.True(t, s.CheckBackSlash(1, 2))
}

func TestForwardSlash(t *testing.T) {
	inputData := scanner.FileToInput(example)
	s := scanner.NewScannerV2(inputData)
	assert.False(t, s.CheckForwardSlash(1, 1))
	assert.False(t, s.CheckForwardSlash(0, 0))
	assert.True(t, s.CheckForwardSlash(1, 2))
	assert.True(t, s.CheckForwardSlash(7, 8))
	assert.False(t, s.CheckForwardSlash(8, 1))

}
