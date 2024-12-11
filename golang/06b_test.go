package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

type Coord2 [2]int
type CoordAndDir [3]int

const (
	up2    = '^'
	down2  = 'v'
	left2  = '<'
	right2 = '>'
)

var offsets2 = map[rune]Coord2{
	up2:    {-1, 0},
	down2:  {1, 0},
	left2:  {0, -1},
	right2: {0, 1},
}

var directions2 = []rune{up2, right2, down2, left2}

func aoc06b() int {

	grid := strings.Split(ReadFile("../inputs/06.txt"), "\n")

	var sx int
	var sy int
	var sd rune
	for x, cc := range grid {
		for y := range cc {
			if cc[y] == up2 || cc[y] == right2 || cc[y] == down2 || cc[y] == left2 {
				sx = x
				sy = y
				sd = rune(cc[y])
			}
		}
	}

	seen := make(map[Coord2]bool)

	move2(sx, sy, grid, sd, seen)

	total := 0
	for k, _ := range seen {
		seen2 := make(map[CoordAndDir]bool)
		if obs(sx, sy, grid, sd, seen2, k) {
			total++
		}
	}

	return total
}

func obs(sx, sy int, grid []string, sd rune, seen map[CoordAndDir]bool, obsCoord Coord2) bool {
	if v, ok := seen[CoordAndDir{sx, sy, int(sd)}]; ok && v {
		return true
	}

	coord := offsets2[sd]

	if grid[sx][sy] == '#' || (sx == obsCoord[0] && sy == obsCoord[1]) {
		bsx := sx - coord[0] // go back
		bsy := sy - coord[1] // go back
		nd, err := turnRight2(sd)
		if err != nil {
			log.Fatal(err)
		}
		return obs(bsx, bsy, grid, nd, seen, obsCoord)
	}

	if sx == 0 || sx == len(grid)-1 || sy == 0 || sy == len(grid[0])-1 {
		return false
	}

	seen[CoordAndDir{sx, sy, int(sd)}] = true
	return obs(sx+coord[0], sy+coord[1], grid, sd, seen, obsCoord)
}

func move2(sx, sy int, grid []string, sd rune, seen map[Coord2]bool) {
	coord := offsets2[sd]

	if grid[sx][sy] == '#' {
		bsx := sx - coord[0] // go back
		bsy := sy - coord[1] // go back
		nd, err := turnRight2(sd)
		if err != nil {
			log.Fatal(err)
		}
		move2(bsx, bsy, grid, nd, seen)
		return
	}

	if sx == 0 || sx == len(grid)-1 || sy == 0 || sy == len(grid[0])-1 {
		seen[Coord2{sx, sy}] = true
		return
	}

	seen[Coord2{sx, sy}] = true
	move2(sx+coord[0], sy+coord[1], grid, sd, seen)
}

func turnRight2(cd rune) (rune, error) {
	for i, d := range directions2 {
		if cd == d {
			return directions2[(i+1)%len(directions2)], nil
		}
	}

	return 0, fmt.Errorf("could not turn right for %c", cd)
}

func Test06b(t *testing.T) {
	t.Log("ANSWER:", aoc06b())
}
