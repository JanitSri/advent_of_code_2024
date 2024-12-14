package main

import (
	"testing"
)

func aoc08b() int {

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
	for _, l := range m {
		getAntiNodesB(l, &p, r, c)
	}

	return len(p)
}

func getAntiNodesB(l []Coord, p *[]Coord, r, c int) {
	for _, l1 := range l {
		for _, l2 := range l {
			if l1 == l2 {
				continue
			}

			l1x := l1[0]
			l1y := l1[1]

			l2x := l2[0]
			l2y := l2[1]

			nx := (l1x - l2x)
			ny := (l1y - l2y)

			dx := l1x
			dy := l1y

			for dx >= 0 && dx < r && dy >= 0 && dy < c {

				pc := Coord{dx, dy}
				if !Contains(*p, pc) {
					*p = append(*p, pc)
				}

				dx += nx
				dy += ny
			}
		}
	}

}

func Test08b(t *testing.T) {
	t.Log("ANSWER", aoc08b())
}
