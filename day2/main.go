package day2

import (
	_ "embed"
	"strconv"
	"strings"
)

// Find how many safe reports there are
// Example
// 7 6 4 2 1

// Needs to be always increasing/decreasing
// Needs to increment by 1-3 amounts (inclusive)

//go:embed input.txt
var input string

//go:embed example.txt
var example string

func ToNums(report string) []int64 {
	numsStr := strings.Split(report, " ")
	nums := []int64{} // Want to use Array over Slice here to see how working with Arrays are.
	for i := range numsStr {
		num, _ := strconv.ParseInt(numsStr[i], 10, 32)
		nums = append(nums, num)
	}

	return nums
}

func validIncrementing(report []int64) bool {
	// 1 3 5 7 9
	currNum := report[0]
	for i := 1; i < len(report); i++ {
		if currNum >= report[i] || (report[i]-currNum > 3) {
			return false
		}
		currNum = report[i]
	}
	return true
}

func validDecrementing(report []int64) bool {
	// 7 6 4 2 1
	currNum := report[0]
	for i := 1; i < len(report); i++ {
		if currNum <= report[i] || (currNum-report[i] > 3) {
			return false
		}
		currNum = report[i]
	}
	return true
}

func validIncrementing2(report []int64) (bool, int) {
	// 1 3 2 4 5 = safe
	unsafeFound := false
	currNum := report[0]
	for i := 1; i < len(report); i++ {
		if (currNum >= report[i] || (report[i]-currNum > 3)) && unsafeFound {
			return false, i
		}
		if currNum >= report[i] || (report[i]-currNum > 3) {
			unsafeFound = true
		}
		currNum = report[i]
	}
	return true, 0
}

func validDecrementing2(report []int64) bool {
	// 7 6 4 2 1
	unsafeFound := false
	currNum := report[0]
	for i := 1; i < len(report); i++ {
		if (currNum <= report[i] || (currNum-report[i] > 3)) && unsafeFound {
			return false
		}
		if currNum <= report[i] || (currNum-report[i] > 3) {
			return false
		}
		currNum = report[i]
	}
	return true
}

func ValidReport(report []int64, version int) bool {
	if version == 1 {
		if validIncrementing(report) {
			return true
		}
		return validDecrementing(report)
	}
	valid, _ := validIncrementing2(report)
	if valid {
		return true
	}
	return validDecrementing2(report)
}

func Main(version int) int {
	validReports := 0
	reports := strings.SplitSeq(input, "\n")
	for report := range reports {
		if ValidReport(ToNums(report), version) {
			validReports++
		}
	}
	return validReports
}

func Example(version int) int {
	validReports := 0
	reports := strings.SplitSeq(example, "\n")
	for report := range reports {
		numReport := ToNums(report)
		if ValidReport(numReport, version) {
			validReports++
		}

	}
	return validReports
}
