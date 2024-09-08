package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("input/day-1.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(file, []byte("\n"))
	for _, v := range lines {
		fmt.Println(v)
	}
}
