package main

import (
	"testing"
)

func aoc08a() int {

	grid, r, c := GetGrid("../inputs/08.txt")

	m := make(map[rune][]Coord)
	for rr := range r {
		for cc := range c {
			if rune(grid[rr][cc]) != '.' {
				l := m[rune(grid[rr][cc])]
				m[rune(grid[rr][cc])] = append(l, Coord{rr, cc})
			}
		}
	}

	p := []Coord{}
	for k, l := range m {
		getAntiNodes(l, &p, grid, r, c, k)
	}

	return len(p)
}

func getAntiNodes(l []Coord, p *[]Coord, grid []string, r, c int, k rune) {
	for _, l1 := range l {
		for _, l2 := range l {
			if l1 == l2 {
				continue
			}

			l1x := l1[0]
			l1y := l1[1]

			l2x := l2[0]
			l2y := l2[1]

			dx := l1x + (l1x - l2x)
			dy := l1y + (l1y - l2y)

			pc := Coord{dx, dy}
			if dx < 0 || dx >= r || dy < 0 || dy >= c || rune(grid[dx][dy]) == k || Contains(*p, pc) {
				continue
			}

			*p = append(*p, pc)
		}
	}

}

func Test08a(t *testing.T) {
	t.Log("ANSWER", aoc08a())
}
