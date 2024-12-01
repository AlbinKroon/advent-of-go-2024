package part2

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

func Solve() string {
	var result int64
	left, right := parse()
	for i := 0; i < len(left); i++ {
		var similarity int64
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				similarity++
			}
		}
		result += int64(left[i]) * similarity
		similarity = 0
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
	return left, right
}
