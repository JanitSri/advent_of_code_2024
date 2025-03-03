package main

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func aoc03a() int {
	c := ReadFile("../inputs/03.txt")

	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	ll := r.FindAll([]byte(c), -1)

	total := 0
	for _, l := range ll {
		total += multiply(string(l))
	}

	return total
}

func multiply(input string) int {
	i := strings.Split(input, ",")
	left, _ := strconv.Atoi(strings.TrimPrefix(i[0], "mul("))
	right, _ := strconv.Atoi(strings.TrimSuffix(i[1], ")"))

	return left * right
}

func Test03a(t *testing.T) {
	t.Log("ANSWER:", aoc03a())
}
