package main

import (
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
)

type stoneTask struct {
	iter int
	num  int
}

func aoc11a() int32 {

	ls := strings.Split(ReadFile("../inputs/11.txt"), " ")

	nums := []int{}
	for _, l := range ls {
		n, _ := strconv.Atoi(l)
		nums = append(nums, n)
	}

	tasks := make(chan *stoneTask, 100)

	for _, n := range nums {
		tasks <- &stoneTask{iter: 0, num: n}
	}

	var wg sync.WaitGroup
	go func() {
		wg.Wait()
		close(tasks)
	}()

	numOfIterations := 25
	var r int32
	for task := range tasks {
		wg.Add(1)
		go func(st *stoneTask) {
			defer wg.Done()
			ns := stoneOperation(st)
			for _, s := range ns {
				if s.iter == numOfIterations {
					atomic.AddInt32(&r, 1)
				} else {
					tasks <- s
				}
			}
			return
		}(task)
	}

	return r
}

func stoneOperation(st *stoneTask) []*stoneTask {
	i := st.iter + 1
	n := st.num
	if n == 0 {
		return []*stoneTask{{iter: i, num: 1}}
	} else if CountDigits(n)%2 == 0 {
		r := SplitDigitsInHalf(n)
		return []*stoneTask{{iter: i, num: r[0]}, {iter: i, num: r[1]}}
	}

	return []*stoneTask{{iter: i, num: n * 2024}}
}

func Test11a(t *testing.T) {
	t.Log("ANSWER", aoc11a())
}
