package main

import (
	"fmt"
	"strings"
	"testing"
)

func aoc04b() int {
	contents := strings.Split(ReadFile("../inputs/04a.txt"), "\n")

	ws := make([][]rune, len(contents))
	for i, str := range contents {
		ws[i] = []rune(str)
	}

	total := 0
	for ri, r := range ws {
		for ci, _ := range r {
			if ws[ri][ci] == 'A' {
				d1 := (traverse2(ri-1, ci-1, 'M', ws) && traverse2(ri+1, ci+1, 'S', ws)) || traverse2(ri-1, ci-1, 'S', ws) && traverse2(ri+1, ci+1, 'M', ws)
				d2 := (traverse2(ri-1, ci+1, 'M', ws) && traverse2(ri+1, ci-1, 'S', ws)) || (traverse2(ri-1, ci+1, 'S', ws) && traverse2(ri+1, ci-1, 'M', ws))

				if d1 && d2 {
					total++
				}
			}
		}
	}

	return total
}

func traverse2(ri, ci int, r rune, arr [][]rune) bool {
	if ri < 0 || ri >= len(arr) || ci < 0 || ci >= len(arr[ri]) {
		return false
	}

	return arr[ri][ci] == r
}

func Test04b(t *testing.T) {
	fmt.Println("ANSWER:", aoc04b())
}
