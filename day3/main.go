package day3

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

var validChars = []string{"m", "u", "l", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ",", "(", ")"}

func isValid(char string) bool {
	return slices.Contains(validChars, char)
}

func StartBuffering(inputData string) []string {
	validBuff := false
	operations := []string{}
	buffer := ""
	strData := strings.SplitSeq(inputData, "")

	for char := range strData {
		if char == "m" {
			buffer = ""
			validBuff = true
		}
		validBuff = validBuff && isValid(char)
		if !validBuff {
			buffer = ""
			continue
		}
		buffer += char
		if len(buffer) == 4 {
			if buffer != "mul(" {
				validBuff = false
			}
		}
		if char == ")" {
			if !strings.Contains(buffer, ",") {
				buffer = ""
				continue
			}
			if len(buffer) > 7 {
				operations = append(operations, buffer)
				buffer = ""
			}
		}
	}

	return operations
}

func calculateOperations(operations []string) (int, error) {
	var total int64
	for _, op := range operations {
		parts := strings.Split(op, ",")
		if len(parts) != 2 {
			return 0, fmt.Errorf("operation parts was not 2: %s", parts)
		}
		if !strings.HasPrefix(parts[0], "mul(") {
			return 0, fmt.Errorf("part0 did not contain mul(: %s", parts[0])
		}
		if !strings.HasSuffix(parts[1], ")") {
			return 0, fmt.Errorf("part1 did not contain ): %s", parts[1])
		}
		num1 := strings.TrimPrefix(parts[0], "mul(")
		num2 := strings.TrimSuffix(parts[1], ")")

		int1, err := strconv.ParseInt(num1, 10, 32)
		if err != nil {
			return 0, err
		}
		int2, err := strconv.ParseInt(num2, 10, 32)
		if err != nil {
			return 0, err
		}
		total += int1 * int2
	}
	return int(total), nil
}

func Example(puzzle int) int {
	operations := StartBuffering(example)

	fmt.Println(operations)

	total, err := calculateOperations(operations)
	if err != nil {
		fmt.Println(err.Error())
	}

	return total
}

func Main(puzzle int) int {
	operations := StartBuffering(input)

	fmt.Println(operations)

	total, err := calculateOperations(operations)
	if err != nil {
		fmt.Println(err.Error())
	}

	return total
}
