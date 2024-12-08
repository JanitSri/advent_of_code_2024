package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func aoc01b() {
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

	m := make(map[int]int)
	for _, x := range second {
		if val, ok := m[x]; !ok {
			m[x] = 1
		} else {
			m[x] = 1 + val
		}
	}

	total := 0
	for _, x := range first {
		if val, ok := m[x]; ok {
			total += x * val
		}
	}

	fmt.Println("ANSWER:", total)
}

func Test01b(t *testing.T) {
	aoc01b()
}
