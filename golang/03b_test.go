package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func aoc03b() {
	content := ReadFile("../inputs/03.txt")

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllStringIndex(content, -1)

	total := 0
	fullStart := 0
	enabled := true
	for _, m := range matches {
		start := m[0]
		before := content[fullStart:start]

		d := strings.LastIndex(before, "do()")
		dd := strings.LastIndex(before, "don't()")

		if dd > d {
			enabled = false
			fullStart = dd + len("don't()")
		} else if d > dd {
			enabled = true
			fullStart = d + len("do()")
		}

		if enabled {
			total += multiply2(content[m[0]:m[1]])
		}
	}

	fmt.Println(total)
}

func multiply2(input string) int {
	i := strings.Split(input, ",")
	left, _ := strconv.Atoi(strings.TrimPrefix(i[0], "mul("))
	right, _ := strconv.Atoi(strings.TrimSuffix(i[1], ")"))

	return left * right
}

func Test03b(t *testing.T) {
	aoc03b()
}
