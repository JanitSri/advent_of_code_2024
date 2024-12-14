package main

import (
	"fmt"
	"strconv"
	"testing"
)

const (
	DotRune2 string = "."
)

func aoc09b() int {
	c := ReadFile("../inputs/09.txt")

	m := diskMap2(c)
	// displayDiskMap2(m)

	for i := len(m) - 1; i > 0; {
		curr := m[i]

		if curr == DotRune2 {
			i--
			continue
		}

		cnt := 0
		si := i
		for i > 0 && m[i] == curr {
			cnt++
			i--
		}

		done := false
		for j := 0; j < si && !done; j++ {
			if m[j] == DotRune2 {
				sj := j

				cntj := 0
				for j < si && m[j] == DotRune2 {
					cntj++
					j++
				}

				if cnt <= cntj {
					for k := 0; k < cnt; k++ {
						m[sj] = curr
						m[si] = DotRune2

						si--
						sj++
					}
					done = true

				}
			}
		}

	}

	// displayDiskMap2(m)
	return calculateChecksum2(m)
}

func checkSpace(m []string) map[int][]int {
	dm := make(map[int][]int)
	i := 0
	for i < len(m) {
		cnt := 0
		s := i
		for i < len(m) && m[i] == DotRune2 {
			cnt++
			i++
		}

		if cnt > 0 {
			dm[cnt] = append(dm[cnt], s)
		}
		i++
	}

	return dm
}

func diskMap2(c string) []string {

	i := 0
	l := []string{}
	for x, n := range c {
		rr := DotRune2
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

func displayDiskMap2(r []string) {

	for _, rr := range r {
		fmt.Print(rr)
	}
	fmt.Println()
}

func calculateChecksum2(m []string) int {

	total := 0
	for i, rr := range m {
		if rr == DotRune2 {
			continue
		}

		n, _ := strconv.Atoi(rr)
		total += i * n
	}

	return total
}

func Test09b(t *testing.T) {
	t.Log("ANSWER", aoc09b())
}
