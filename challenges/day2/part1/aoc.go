package part1

import (
	_ "embed"
	"math"
	"slices"
	"strconv"
	"strings"
)

//go:embed input
var input string

func Solve() string {
	var result int
	reports := parse()
	for i := 0; i < len(reports); i++ {
		isSafe := checkIfSafe(reports[i])
		if isSafe {
			result++
		}
	}
	return strconv.Itoa(result)
}

func checkIfSafe(report []int) bool {
	isSorted := checkIfSorted(report)
	if isSorted {
		return checkIfSafeDistance(report)
	}
	return false
}

func checkIfSafeDistance(report []int) bool {
	prevHeight := -1
	for _, v := range report {
		if prevHeight == -1 {
			prevHeight = v
			continue
		}
		diff := math.Abs(float64(prevHeight - v))
		if diff < 1 || diff > 3 {
			return false
		}
		prevHeight = v
	}
	return true
}

func checkIfSorted(report []int) bool {
	if slices.IsSorted(report) {
		return true
	}
	sorted := slices.Clone(report)
	slices.Sort(sorted)
	slices.Reverse(sorted)

	return slices.Equal(sorted, report)
}

func parse() [][]int {
	rows := strings.Split(input, "\n")
	result := make([][]int, len(rows))
	for i, row := range rows {
		row = strings.TrimSpace(row)
		split := strings.Split(row, " ")

		parsedRow := make([]int, len(split))

		for j, v := range split {
			parsedRow[j], _ = strconv.Atoi(v)
		}
		result[i] = parsedRow
	}
	return result
}
