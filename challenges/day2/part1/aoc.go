package part1

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed input
var input string

func Solve() string {
	var result int64
	// TODO := parse()
	for i := 0; i < len(TODO); i++ {
		// TODO
	}
	return strconv.FormatInt(result, 10)
}

func parse() () {
	rows := strings.Split(input, "\n")
	for i, row := range rows {
		row = strings.TrimSpace(row)
		split := strings.Split(row, "")
	}
	return
}
