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

//go:embed example2.txt
var example2 string

var validChars = []string{"m", "u", "l", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ",", "(", ")"}

func isValid(char string) bool {
	return slices.Contains(validChars, char)
}

func StartBuffering(inputData string) []string {
	validBuff := false
	operations := &[]string{}
	var buffer *string
	strData := strings.SplitSeq(inputData, "")

	for char := range strData {
		buffer, validBuff = updateBuffer(char, buffer, validBuff, operations)
	}

	return *operations
}

func BufferingWithDoDont(inputData string) []string {
	operations := &[]string{}

	active := true
	var buffer *string
	validBuffer := true

	array := strings.Split(inputData, "")
	for i := 0; i < len(array); i++ {
		char := array[i]
		if char == "d" {
			do := strings.Join(array[i:i+4], "")
			dont := strings.Join(array[i:i+7], "")

			if do == "do()" {
				active = true
				i += 3
				continue
			}

			if dont == "don't()" {
				active = false
				i += 6
				continue
			}
		}

		if !active {
			continue
		}

		buffer, validBuffer = updateBuffer(char, buffer, validBuffer, operations)
	}

	return *operations
}

func updateBuffer(char string, buffer *string, validBuff bool, operations *[]string) (*string, bool) {
	emptyStr := new(string)
	if char == "m" {
		buffer = emptyStr
		validBuff = true
	}
	validBuff = validBuff && isValid(char)
	if !validBuff {
		buffer = emptyStr
		return buffer, validBuff
	}
	*buffer += char
	if len(*buffer) == 4 {
		if *buffer != "mul(" {
			validBuff = false
		}
	}
	if char == ")" {
		if !strings.Contains(*buffer, ",") {
			buffer = emptyStr
			return buffer, validBuff
		}
		if len(*buffer) > 7 {
			*operations = append(*operations, *buffer)
			buffer = emptyStr
		}
	}
	return buffer, validBuff
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
	var operations []string
	if puzzle == 1 {
		operations = StartBuffering(example)
	} else {
		operations = StartBuffering(example2)
	}

	fmt.Println(operations)

	total, err := calculateOperations(operations)
	if err != nil {
		fmt.Println(err.Error())
	}

	return total
}

func Main(puzzle int) int {
	var operations []string
	if puzzle == 1 {
		operations = StartBuffering(input)
	} else {
		operations = BufferingWithDoDont(input)
	}

	total, err := calculateOperations(operations)
	if err != nil {
		fmt.Println(err.Error())
	}

	return total
}
