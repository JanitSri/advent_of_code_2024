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

	return r
}

func stoneOperationB(st *stoneTaskB) []*stoneTaskB {
	i := st.iter + 1
	n := st.num
	if n == 0 {
		return []*stoneTaskB{{iter: i, num: 1}}
	} else if CountDigits(n)%2 == 0 {
		r := SplitDigitsInHalf(n)
		return []*stoneTaskB{{iter: i, num: r[0]}, {iter: i, num: r[1]}}
	}

	return []*stoneTaskB{{iter: i, num: n * 2024}}
}

func Test11b(t *testing.T) {
	t.Log("ANSWER", aoc11b())
}
