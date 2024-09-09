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

	total := 0
	contents := day1.BagContents{Blue: 14, Green: 13, Red: 12}
	lines := strings.Split(string(file), "\n")
	for _, v := range lines {
		game, err := day1.NewGame(v)
		if err != nil {
			log.Fatal(err)
		}

		if game.IsPossible(contents) {
			total += game.Id
		}
	}

	fmt.Printf("total is: %d", total)
}
