package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

const (
	up    = '^'
	down  = 'v'
	left  = '<'
	right = '>'
)

var offsets = map[rune]Coord{
	up:    {-1, 0},
	down:  {1, 0},
	left:  {0, -1},
	right: {0, 1},
}

var directions = []rune{up, right, down, left}

func aoc06a() int {

	grid := strings.Split(ReadFile("../inputs/06.txt"), "\n")

	var sx int
	var sy int
	var sd rune
	for x, cc := range grid {
		for y := range cc {
			if cc[y] == up || cc[y] == right || cc[y] == down || cc[y] == left {
				sx = x
				sy = y
				sd = rune(cc[y])
			}
		}
	}

	seen := make(map[Coord]bool)

	move(sx, sy, grid, sd, seen)

	return len(seen)
}

func move(sx, sy int, grid []string, sd rune, seen map[Coord]bool) {
	coord := offsets[sd]

	if grid[sx][sy] == '#' {
		bsx := sx - coord[0] // go back
		bsy := sy - coord[1] // go back
		nd, err := turnRight(sd)
		if err != nil {
			log.Fatal(err)
		}
		move(bsx, bsy, grid, nd, seen)
		return
	}

	if sx == 0 || sx == len(grid)-1 || sy == 0 || sy == len(grid[0])-1 {
		seen[Coord{sx, sy}] = true
		return
	}

	seen[Coord{sx, sy}] = true
	move(sx+coord[0], sy+coord[1], grid, sd, seen)
}

func turnRight(cd rune) (rune, error) {
	for i, d := range directions {
		if cd == d {
			return directions[(i+1)%len(directions)], nil
		}
	}

	return 0, fmt.Errorf("could not turn right for %c", cd)
}

func Test06a(t *testing.T) {
	t.Log("ANSWER:", aoc06a())
}
