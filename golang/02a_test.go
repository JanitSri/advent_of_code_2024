package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func aoc02a() int {
	f, err := os.Open("../inputs/02.txt")
	if err != nil {
		log.Fatal("cannot open file:", err)
	}

	r := bufio.NewReader(f)

	total := 0
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal("could not read file", err)
		}

		if safe2a(l) {
			total++
		}
	}

	return total
}

func safe2a(line string) bool {
	items := strings.Split(strings.TrimSuffix(line, "\n"), " ")

	allInc := true
	allDec := true
	for i := 1; i < len(items); i++ {
		curr, _ := strconv.Atoi(items[i])
		prev, _ := strconv.Atoi(items[i-1])

		allInc = allInc && curr > prev
		allDec = allDec && curr < prev

		diff := math.Abs(float64(curr - prev))
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return allInc || allDec
}

func Test02a(t *testing.T) {
	t.Log("ANSWER:", aoc02a())
}
