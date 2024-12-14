package main

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	DotRune string = "."
)

func aoc09a() int {

	c := ReadFile("../inputs/09.txt")

	m := diskMap(c)
	// DisplayDiskMap(m)

	l := 0
	r := len(m) - 1

	for l < r {
		lc := m[l]
		rc := m[r]

		if lc == DotRune {
			m[l] = rc
			m[r] = DotRune

			for r > l && m[r] == DotRune {
				r -= 1
			}
		}
		l += 1
		// DisplayDiskMap(m)
	}

	return calculateChecksum(m)
}

func diskMap(c string) []string {

	i := 0
	l := []string{}
	for x, n := range c {
		rr := DotRune
		if x%2 == 0 {
			rr = strconv.Itoa(i)
			i += 1
		}

		nn, _ := strconv.Atoi(string(n))
		for nn > 0 {
			l = append(l, rr)
			nn -= 1
		}
	}

	return l
}

func displayDiskMap(r []string) {

	for _, rr := range r {
		fmt.Print(rr)
	}
	fmt.Println()
}

func calculateChecksum(m []string) int {

	total := 0
	for i, rr := range m {
		if rr == DotRune {
			continue
		}

		n, _ := strconv.Atoi(rr)
		total += i * n
	}

	return total
}

func Test09a(t *testing.T) {
	t.Log("ANSWER", aoc09a())
}
