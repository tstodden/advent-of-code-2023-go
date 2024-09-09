package main

import (
	"aoc2023/pkg/day1"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("input/day-1.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")
	for _, v := range lines {
		game, err := day1.NewGame(v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", game)
	}
}
