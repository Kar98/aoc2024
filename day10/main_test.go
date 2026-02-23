package day10

import (
	"aoc/common"
	"aoc/day10/hiker"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed eg1.txt
var example1 string

//go:embed eg2.txt
var example2 string

//go:embed eg3.txt
var example3 string

//go:embed eg4.txt
var example4 string

func TestFileToInput(t *testing.T) {
	file := hiker.FileToInput(example1)
	assert.Len(t, file, 10)
}

func TestGetStartingPoints(t *testing.T) {
	hiker := hiker.NewHiker(example1)
	points := hiker.GetStartingPoints()
	assert.Len(t, points, 1)
	assert.Equal(t, 0, points[0].R)
}

func TestHikeExample1(t *testing.T) {
	total := []hiker.Coordinate{}
	hike := hiker.NewHiker(example1)
	points := hike.GetStartingPoints()
	total, err := hike.Hike(points[0], total)

	assert.NoError(t, err)
	assert.Len(t, total, 4)

}

func TestHikeExample2(t *testing.T) {
	total := []hiker.Coordinate{}
	hike2 := hiker.NewHiker(example2)
	points := hike2.GetStartingPoints()
	for _, point := range points {
		total, _ = hike2.Hike(point, total)
	}
	t.Log(total)
	total = common.GetUniques(total)
	assert.Len(t, total, 4)
}

func TestHikeExample3(t *testing.T) {
	total := []hiker.Coordinate{}
	hike2 := hiker.NewHiker(example3)
	points := hike2.GetStartingPoints()

	score := 0
	for _, point := range points {
		endPoints, _ := hike2.Hike(point, []hiker.Coordinate{})
		total = common.GetUniques(endPoints)
		score += len(total)
	}
	assert.Equal(t, 3, score)
}

func TestHikeExample4(t *testing.T) {
	total := []hiker.Coordinate{}
	hike2 := hiker.NewHiker(example4)
	points := hike2.GetStartingPoints()

	score := 0
	for _, point := range points {
		endPoints, _ := hike2.Hike(point, []hiker.Coordinate{})
		total = common.GetUniques(endPoints)
		score += len(total)
	}
	assert.Equal(t, 36, score)
}
