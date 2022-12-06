package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt", 14))

}

func Challenge(fileName string, markerLevel int) int {
	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	input := fileScanner.Text()

	queue := []rune(input[0:markerLevel])
	index := markerLevel

	for true {
		containsDouble, whichIndex := containsDouble(queue)
		if !containsDouble {
			return index
		}
		queue = append(queue[whichIndex:], []rune(input[index:index+whichIndex])...)

		index += whichIndex
		if index >= len(input) {
			break
		}
	}

	return -1
}

func containsDouble(queue []rune) (bool, int) {
	littleHisto := make(map[rune]int)
	for index, rne := range queue {
		if _, exists := littleHisto[rne]; exists {
			return true, littleHisto[rne] + 1
		}
		littleHisto[rne] = index
	}
	return false, 0
}
