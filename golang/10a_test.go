package main

import (
	"strconv"
	"testing"
)

func aoc10a() int {

	grid, r, c := GetGrid("../inputs/10.txt")

	p := []Coord{}
	g := [][]int{}
	for rr := range r {
		g = append(g, []int{})
		for cc := range c {
			if grid[rr][cc] == '.' {
				g[rr] = append(g[rr], -1)
				continue
			}
			n, _ := strconv.Atoi(string(grid[rr][cc]))
			g[rr] = append(g[rr], n)
			if g[rr][cc] == 0 {
				p = append(p, Coord{rr, cc})
			}
		}
	}

	total := 0
	for _, pp := range p {
		total += findTrail(pp[0], pp[1], 0, g)

		ResetGrid(g, -1, 9)
	}

	return total
}

func findTrail(x, y, curr int, g [][]int) int {
	if x < 0 || x >= len(g) || y < 0 || y >= len(g[0]) || g[x][y] != curr {
		return 0
	}

	if curr == 9 && g[x][y] == curr {
		g[x][y] = -1
		return 1
	}

	return findTrail(x+DirectionMap["up"][0], y+DirectionMap["up"][1], curr+1, g) +
		findTrail(x+DirectionMap["down"][0], y+DirectionMap["down"][1], curr+1, g) +
		findTrail(x+DirectionMap["left"][0], y+DirectionMap["left"][1], curr+1, g) +
		findTrail(x+DirectionMap["right"][0], y+DirectionMap["right"][1], curr+1, g)
}

func Test10a(t *testing.T) {
	t.Log("ANSWER", aoc10a())
}
