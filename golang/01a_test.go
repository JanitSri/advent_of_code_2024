package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func aco01a() int {
	input := ReadFile("../inputs/01.txt")
	lines := strings.Split(input, "\n")

	first := []int{}
	second := []int{}
	for _, line := range lines {
		l := strings.Split(line, "   ")

		f, _ := strconv.Atoi(l[0])
		s, _ := strconv.Atoi(l[1])
		first = append(first, f)
		second = append(second, s)
	}

	sort.IntSlice(first).Sort()
	sort.IntSlice(second).Sort()

	total := 0
	for i := range first {
		diff := first[i] - second[i]
		total += int(math.Abs(float64(diff)))
	}

	return total
}

func Test01a(t *testing.T) {
	t.Log("ANSWER:", aco01a())
}
