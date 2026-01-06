package main

import (
	"aoc/day4"
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	output := day4.Main(1)
	duration := time.Since(start)
	fmt.Println("Part1")
	fmt.Println(output)
	fmt.Printf("time taken in us: %d\n", duration.Microseconds())

	output2 := day4.Main(2)
	fmt.Println("Part2")
	fmt.Println(output2)

	// output = day4.Main(2)
	// fmt.Println("Part2")
	// fmt.Println(output)
}
