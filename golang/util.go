package main

import (
	"log"
	"os"
)

func ReadFile(path string) string {
	b, err := os.ReadFile(path)

	if err != nil {
		log.Fatalln("could not read file", path)
	}

	return string(b)
}
