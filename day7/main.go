package day7

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

func FileToInput(file string) ([][]int64, error) {
	lines := strings.Split(file, "\n")
	output := make([][]int64, len(lines))
	for i, line := range lines {
		replaced := strings.ReplaceAll(line, ":", "")
		nums := strings.Split(replaced, " ")
		tmp := []int64{}
		for _, num := range nums {
			iNum, err := strconv.ParseInt(num, 10, 64)
			if err != nil {
				return output, err
			}
			tmp = append(tmp, iNum)
		}
		output[i] = tmp
	}

	return output, nil
}

func Main(part int) int64 {
	rows, err := FileToInput(input)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	if part == 1 {
		var total int64
		for _, row := range rows {
			if isValidOperation(row) {
				total += row[0]
			}
		}
		return total
	}

	return 0
}

func isValidOperation(nums []int64) bool {
	main := nums[0]
	rightSideNumbers := nums[1:]
	operators := generateOperators(rightSideNumbers)
	for _, listOfSymbols := range operators {
		runningTotal := rightSideNumbers[0]
		// listOfSymbols = x + x
		for i := 0; i < len(rightSideNumbers)-1; i++ {
			if string(listOfSymbols[i]) == "0" {
				runningTotal *= rightSideNumbers[i+1]
			} else {
				runningTotal += rightSideNumbers[i+1]
			}
		}
		if main == runningTotal {
			return true
		}
	}

	return false
}

func generateOperators(nums []int64) []string {
	// n^2 - 1
	if len(nums) == 2 {
		return []string{"0", "1"}
	}
	if len(nums) == 3 {
		return []string{"00", "01", "10", "11"}
	}
	n := len(nums) - 1
	totalOperators := n*n - 1
	operators := []string{}
	for i := range totalOperators {
		asBin := fmt.Sprintf("%b", i)
		operators = append(operators, pad(asBin, n))
	}

	return operators
}

func pad(input string, maxPad int) string {
	len := len(input)
	if maxPad-len <= 0 {
		return input
	}
	return strings.Repeat("0", maxPad-len) + input
}

func convertBinToOperators(binaryNumbers []string) []string {
	operators := []string{}
	for _, binNum := range binaryNumbers {
		operator := ""
		nums := strings.Split(binNum, "")
		for _, num := range nums {
			if num == "0" {
				operator += "x"
			} else {
				operator += "+"
			}
		}
		operators = append(operators, operator)
	}
	return operators
}
