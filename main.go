package main

import (
	"aoc2023/pkg/day2"
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
	lines := strings.Split(string(file), "\n")
	for _, v := range lines {
		game, err := day2.NewGame(v)
		if err != nil {
			log.Fatal(err)
		}

		contents := game.GetMinimumBagContents()
		total += contents.Blue * contents.Green * contents.Red
	}

	fmt.Printf("total is: %d", total)
}
