package main

import (
	"fmt"
	"strings"
	"testing"
)

var (
	xc = []rune{'X', 'M', 'A', 'S'}
)

func aoc04a() int {
	contents := strings.Split(ReadFile("../inputs/04.txt"), "\n")

	ws := make([][]rune, len(contents))
	for i, str := range contents {
		ws[i] = []rune(str)
	}

	total := 0
	directions := [][2]int{
		{-1, 0},  // Up
		{-1, 1},  // Up-Right
		{0, 1},   // Right
		{1, 1},   // Down-Right
		{1, 0},   // Down
		{1, -1},  // Down-Left
		{0, -1},  // Left
		{-1, -1}, // Up-Left
	}
	for ri, r := range ws {
		for ci := range r {
			for _, d := range directions {
				if traverse(ri, ci, 0, ws, d) {
					total++
				}
			}
		}
	}

	return total
}

func traverse(ri, ci, xi int, arr [][]rune, d [2]int) bool {
	if ri < 0 || ri >= len(arr) || ci < 0 || ci >= len(arr[ri]) || xi >= len(xc) {
		return false
	}

	if arr[ri][ci] == xc[xi] {
		if xi == len(xc)-1 {
			return true
		}

		return traverse(ri+d[0], ci+d[1], xi+1, arr, d)
	}

	return false
}

func Test04a(t *testing.T) {
	fmt.Println("ANSWER:", aoc04a())
}
