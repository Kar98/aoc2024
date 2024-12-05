package main

import (
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fileData, err := os.ReadFile("input.txt")
	if err != nil {
		slog.Error("error", "e", err.Error())
		return
	}
	file := string(fileData)
	lines := strings.Split(file, "\n")
	var left []int
	var right []int
	for _, line := range lines {
		splits := strings.Split(line, "   ")
		left = append(left, toint(splits[0]))
		right = append(right, toint(splits[1]))
	}
	pbm1 := problem1(left, right)
	pbm2 := problem2(left, right)

	fmt.Println("problem 1: ", pbm1)
	fmt.Println("problem 2: ", pbm2)
}

func problem1(left, right []int) int {

	slices.Sort(left)
	slices.Sort(right)

	total := 0

	for i := 0; i < len(left); i++ {
		total += abs(left[i] - right[i])
	}
	return total
}

func problem2(left, right []int) int {
	similarityScore := 0
	for _, num := range left {
		found := count(num, right)
		similarityScore += num * found
	}
	return similarityScore
}

func count(num int, right []int) int {
	found := 0
	for _, rightNum := range right {
		if num == rightNum {
			found++
		}
	}
	return found
}

func toint(num string) int {
	parse, err := strconv.ParseInt(num, 10, 32)
	if err != nil {
		slog.Error("error parsing")
	}
	return int(parse)
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}
