package part1

import (
	_ "embed"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input
var input string

func Solve() string {
	var result int64
	left, right := parse()
	for i := 0; i < len(left); i++ {
		result += int64(math.Abs(float64(left[i] - right[i])))
	}
	return strconv.FormatInt(result, 10)
}

func parse() ([]int, []int) {
	rows := strings.Split(input, "\n")
	left := make([]int, len(rows))
	right := make([]int, len(rows))
	for i, row := range rows {
		row = strings.TrimSpace(row)
		lr := strings.Split(row, "   ")
		left[i], _ = strconv.Atoi(lr[0])
		right[i], _ = strconv.Atoi(lr[1])
	}
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})
	return left, right
}
