package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"testing"
)

func aoc07b() int {

	c := strings.Split(ReadFile("../inputs/07.txt"), "\n")

	total := 0
	for _, l := range c {
		ll := strings.Split(l, ":")
		t, _ := strconv.Atoi(ll[0])
		li := strings.Split(strings.TrimSpace(ll[1]), " ")

		btt := make([]int, int((math.Pow(3, float64(len(li)))-1)/2))
		fffff, _ := strconv.Atoi(li[0])
		btt[0] = fffff
		btti := 1

		found := false
		for iii := 1; iii < len(li) && !found; iii++ {
			x := li[iii]
			xx, _ := strconv.Atoi(x)
			for jjj := 0; jjj < int(math.Pow(3, float64(iii))); jjj++ {
				pidx := getParentIdxB(btti)
				c := calculateB(btt[pidx], jjj, xx)
				if iii == len(li)-1 && c == t {
					total += t
					found = true
					break
				}
				btt[btti] = c
				btti++
			}
		}
	}

	return total
}

func getParentIdxB(i int) int {
	// (idx−1)/3
	return int(math.Max(0, float64((i-1)/3)))
}

func calculateB(pv, ci, v int) int {
	mod := ci % 3
	switch mod {
	case 0: // Left child
		return pv + v
	case 1: // Middle child
		return pv * v
	case 2: // Right child
		f := strconv.Itoa(pv)
		s := strconv.Itoa(v)
		rrrrr, _ := strconv.Atoi(f + s)
		return rrrrr
	}

	log.Fatal("how the f***k did it get here!!!!! 🤕🫥🥵💀💩🤡👹")
	return 0
}

func Test07b(t *testing.T) {
	t.Log("ANSWER", aoc07b())
}
