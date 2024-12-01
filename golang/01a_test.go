package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func AOC01a() {
	input := ReadFile("../inputs/01a.txt")
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
	for i, _ := range first {
		diff := first[i] - second[i]
		total += int(math.Abs(float64(diff)))
	}

	fmt.Println("ANSWER:", total)
}

func Test01a(t *testing.T) {
	AOC01a()
}
