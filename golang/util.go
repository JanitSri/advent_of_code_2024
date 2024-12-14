package main

import (
	"log"
	"os"
	"strings"
)

type Coord [2]int

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

func Contains[T comparable](a []T, x T) bool {
	for _, i := range a {
		if i == x {
			return true
		}
	}

	return false
}
