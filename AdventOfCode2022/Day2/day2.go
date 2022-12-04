package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt"))
}

func Challenge(fileName string) (int, int) {
	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	guide := initializeVictoryGuide()
	treacheryGuide := initializeTreacheryGuide()

	finalScoreWithoutRigging := 0
	finalScoreWithRigging := 0

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		guideline := strings.Split(str, " ")
		if len(guideline) < 2 {
			continue
		}
		finalScoreWithoutRigging += guide[guideline[1]][guideline[0]]
		finalScoreWithRigging += treacheryGuide[guideline[1]][guideline[0]]
	}

	return finalScoreWithoutRigging, finalScoreWithRigging
}

func initializeVictoryGuide() map[string]map[string]int {
	guide := map[string]map[string]int{}

	guide["X"] = make(map[string]int)
	guide["X"]["A"] = 4
	guide["X"]["B"] = 1
	guide["X"]["C"] = 7

	guide["Y"] = make(map[string]int)
	guide["Y"]["A"] = 8
	guide["Y"]["B"] = 5
	guide["Y"]["C"] = 2

	guide["Z"] = make(map[string]int)
	guide["Z"]["A"] = 3
	guide["Z"]["B"] = 9
	guide["Z"]["C"] = 6

	return guide
}

func initializeTreacheryGuide() map[string]map[string]int {
	guide := map[string]map[string]int{}

	guide["X"] = make(map[string]int)
	guide["X"]["A"] = 3
	guide["X"]["B"] = 1
	guide["X"]["C"] = 2

	guide["Y"] = make(map[string]int)
	guide["Y"]["A"] = 4
	guide["Y"]["B"] = 5
	guide["Y"]["C"] = 6

	guide["Z"] = make(map[string]int)
	guide["Z"]["A"] = 8
	guide["Z"]["B"] = 9
	guide["Z"]["C"] = 7

	return guide
}
