package main

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func aoc05b() int {

	co := strings.Split(ReadFile("../inputs/05.txt"), "\n\n")

	c1 := strings.Split(co[0], "\n")
	c2 := strings.Split(co[1], "\n")

	po := make(map[int][]int)
	for _, l := range c1 {
		parts := strings.Split(l, "|")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])

		if _, ok := po[first]; !ok {
			po[first] = []int{}
		}

		po[first] = append(po[first], second)
	}

	var total int
	for _, l := range c2 {
		a := strings.Split(l, ",")
		if !valid2(a, po) {

			sort.Slice(a, func(i, j int) bool {
				ii, _ := strconv.Atoi(a[i])
				jj, _ := strconv.Atoi(a[j])

				pp := po[jj]

				for _, v := range pp {
					if v == ii {
						return true
					}
				}

				return false
			})

			total += mid2(a)
		}
	}

	return total
}

func mid2(a []string) int {
	r, _ := strconv.Atoi(a[0+(len(a)-1)/2])
	return r
}

func valid2(s []string, po map[int][]int) bool {

	seen := make(map[int]bool)

	for _, r := range s {
		rr, _ := strconv.Atoi(r)
		l := po[rr]

		for _, i := range l {
			if _, ok := seen[i]; ok {
				return false
			}
		}

		seen[rr] = true
	}

	return true
}

func Test05b(t *testing.T) {
	t.Log("ANSWER", aoc05b())
}
