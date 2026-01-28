package day9

import (
	"aoc/day9/diskbuilder"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func Main(puzzle int) int64 {
	if puzzle == 1 {
		builder := diskbuilder.NewSimpleDisk(input)
		err := builder.BuildDisk()
		if err != nil {
			fmt.Println(err.Error())
			return 0
		}
		builder.Sort()
		return builder.CalculateChecksum()
	}

	return 0
}
