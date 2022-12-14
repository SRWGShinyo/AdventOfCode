package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coordinate struct {
	XCoord int
	YCoord int
}

func main() {
	fmt.Println(Challenge("./example_input.txt"))
}

func Challenge(fileName string) int {
	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
	}

	return -1
}
