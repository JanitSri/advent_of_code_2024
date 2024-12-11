package main

import (
	"math"
	"strconv"
	"strings"
	"testing"
)

func aoc07a() int {

	c := strings.Split(ReadFile("../inputs/07.txt"), "\n")

	total := 0
	for _, l := range c {
		ll := strings.Split(l, ":")
		t, _ := strconv.Atoi(ll[0])
		li := strings.Split(strings.TrimSpace(ll[1]), " ")

		bt := make([]int, int(math.Pow(2, float64(len(li))))-1)
		bti := 0

		found := false
		for iii := 0; iii < len(li) && !found; iii++ {
			x := li[iii]
			xx, _ := strconv.Atoi(x)
			for jjj := 0; jjj < int(math.Pow(2, float64(iii))); jjj++ {
				pidx := getParentIdx(bti)
				c := calculate(bt[pidx], jjj, xx)
				if iii == len(li)-1 && c == t {
					total += t
					found = true
					break
				}
				bt[bti] = c
				bti++
			}
		}
	}

	return total
}

func getParentIdx(i int) int {
	// odd -> idx/2 - 1
	// even -> idx/2 - 2
	if i%2 == 0 {
		return int(math.Max(0, float64((i-2)/2)))
	}

	return int(math.Max(0, float64((i-1)/2)))
}

func calculate(pv, ci, v int) int {
	// odd -> multiply
	// even -> add
	if ci%2 == 0 {
		return pv + v
	}

	return pv * v
}

func Test07a(t *testing.T) {
	t.Log("ANSWER", aoc07a())
}
