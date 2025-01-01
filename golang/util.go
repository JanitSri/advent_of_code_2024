package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord [2]int

var DirectionMap = map[string][2]int{
	"up":        {-1, 0},
	"upRight":   {-1, 1},
	"right":     {0, 1},
	"downRight": {1, 1},
	"down":      {1, 0},
	"downLeft":  {1, -1},
	"left":      {0, -1},
	"upLeft":    {-1, -1},
}

func ReadFile(path string) string {
	b, err := os.ReadFile(path)

	if err != nil {
		log.Fatalln("could not read file", path)
	}

	return string(b)
}

func GetGrid(path string) ([]string, int, int) {
	b, err := os.ReadFile(path)

	if err != nil {
		log.Fatalln("could not read file", path)
	}

	grid := strings.Split(string(b), "\n")

	return grid, len(grid), len(grid[0])
}

func ResetGrid[T comparable](g [][]T, o, t T) {
	for r := range g {
		for c := range g[r] {
			if g[r][c] == o {
				g[r][c] = t
			}
		}
	}
}

func Contains[T comparable](a []T, x T) bool {
	for _, i := range a {
		if i == x {
			return true
		}
	}

	return false
}

func CountDigits(i int) int {
	var r int

	for i > 0 {
		i /= 10
		r++
	}

	return r
}

func SplitDigitsInHalf(i int) []int {
	s := strconv.Itoa(i)
	n := CountDigits(i) / 2
	l, _ := strconv.Atoi(s[:n])
	r, _ := strconv.Atoi(s[n:])

	return []int{l, r}
}
