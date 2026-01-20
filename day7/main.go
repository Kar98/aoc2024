package day7

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed example.txt
var example string

//go:embed input.txt
var input string

var candidates = [3]string{"+", "x", "|"}

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

	var total int64
	for _, row := range rows {
		if isValidOperationV2(row) {
			total += row[0]
		}
	}
	return total
}

func isValidOperation(nums []int64) bool {
	main := nums[0]
	rightSideNumbers := nums[1:]
	operators := generateOperators(rightSideNumbers)
	for _, listOfSymbols := range operators {
		runningTotal := rightSideNumbers[0]
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

func isValidOperationV2(nums []int64) bool {
	main := nums[0]
	rightSideNumbers := nums[1:]
	operators := generateOperatorsV2(rightSideNumbers)
	for _, listOfSymbols := range operators {
		runningTotal := rightSideNumbers[0]
		for i := 0; i < len(rightSideNumbers)-1; i++ {
			if string(listOfSymbols[i]) == "x" {
				runningTotal *= rightSideNumbers[i+1]
			} else if string(listOfSymbols[i]) == "+" {
				runningTotal += rightSideNumbers[i+1]
			} else {
				total, err := mergeNumbers(runningTotal, rightSideNumbers[i+1])
				if err != nil {
					fmt.Println(err.Error())
					return false
				}
				runningTotal = total
			}
		}
		if main == runningTotal {
			return true
		}
	}
	return false
}

func mergeNumbers(i1, i2 int64) (int64, error) {
	mergedStr := fmt.Sprint(i1) + fmt.Sprint(i2)
	return strconv.ParseInt(mergedStr, 10, 64)
}

func generateOperators(nums []int64) []string {
	// n^2 - 1
	l := len(nums) - 1
	totalOperators := math.Pow(float64(2), float64(l))
	operators := []string{}

	for i := range int(totalOperators) {
		asBin := fmt.Sprintf("%b", i)
		operators = append(operators, pad(asBin, l))
	}

	return operators
}

func generateOperatorsV2(nums []int64) []string {
	setsOfOperators := len(nums) - 1
	totalPatterns := int(math.Pow(float64(3), float64(setsOfOperators)))

	matrix := make([]string, totalPatterns)
	for i := 0; i < totalPatterns; i++ {
		matrix[i] = createSlice(i, setsOfOperators)
	}

	return matrix
}

func createSlice(i int, size int) string {
	slice := make([]string, size)
	for m := range size {
		factor := int(math.Pow(float64(3), float64(size-1-m)))
		slice[m] = get(i / factor)
	}
	return strings.Join(slice, "")
}

func get(i int) string {
	idx := i % 3
	return candidates[idx]
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
