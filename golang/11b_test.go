package main

import (
	"strconv"
	"strings"
	"testing"
)

type stoneTaskB struct {
	iter int
	num  int
}

func aoc11b() int {

	ls := strings.Split(ReadFile("../inputs/11.txt"), " ")

	nums := []int{}
	for _, l := range ls {
		n, _ := strconv.Atoi(l)
		nums = append(nums, n)
	}

	var r int
	iterCnt := 75
	m := make(map[[2]int]int)
	for _, n := range nums {
		r += stoneCalc(n, iterCnt, m)
	}

	return r
}

func stoneCalc(n, iterCnt int, m map[[2]int]int) int {
	if iterCnt == 0 {
		return 1
	}

	if _, ok := m[[2]int{n, iterCnt}]; !ok {
		var r int

		if n == 0 {
			r = stoneCalc(1, iterCnt-1, m)
		} else if CountDigits(n)%2 == 0 {
			n := SplitDigitsInHalf(n)
			r = stoneCalc(n[0], iterCnt-1, m) + stoneCalc(n[1], iterCnt-1, m)
		} else {
			r = stoneCalc(n*2024, iterCnt-1, m)
		}

		m[[2]int{n, iterCnt}] = r
	}

	return m[[2]int{n, iterCnt}]
}

func Test11b(t *testing.T) {
	t.Log("ANSWER", aoc11b())
}
