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

var operatorMatrix map[int][]string

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
	var base int
	operatorMatrix = make(map[int][]string)
	rows, err := FileToInput(input)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	if part == 1 {
		base = 2
	} else {
		base = 3
	}

	var total int64
	for _, row := range rows {
		if isValidOperation(row, base) {
			total += row[0]
		}
	}
	return total
}

func isValidOperation(nums []int64, base int) bool {
	main := nums[0]
	rightSideNumbers := nums[1:]
	operators := getOperators(rightSideNumbers, base)
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

func getOperators(nums []int64, base int) []string {
	lNums := len(nums)
	res, ok := operatorMatrix[lNums]
	if !ok {
		operators := generateOperators(nums, base)
		operatorMatrix[lNums] = operators
		res = operators
	}
	return res
}

func generateOperators(nums []int64, base int) []string {
	setsOfOperators := len(nums) - 1
	totalPatterns := int(math.Pow(float64(base), float64(setsOfOperators)))

	matrix := make([]string, totalPatterns)
	for i := 0; i < totalPatterns; i++ {
		matrix[i] = createSlice(i, setsOfOperators, base)
	}

	return matrix
}

func createSlice(position int, totalColumns int, base int) string {
	slice := make([]string, totalColumns)
	for m := range totalColumns {
		factor := int(math.Pow(float64(base), float64(totalColumns-1-m)))
		slice[m] = get(position/factor, base)
	}
	return strings.Join(slice, "")
}

func get(i int, base int) string {
	idx := i % base
	return candidates[idx]
}
