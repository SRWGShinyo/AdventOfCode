package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Challenge("./example_input.txt"))
}

func Challenge(fileName string) (int, int) {
	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	finalValueFullOverlap := 0
	finalValuePartialOverlap := 0

	for fileScanner.Scan() {
		str := fileScanner.Text()
		sectionsID := strings.FieldsFunc(str, Split)
		sectionsIDAsInt, err := mapStrConv(sectionsID)
		if err != nil {
			return -1, -1
		}

		if sectionsIDAsInt[0] <= sectionsIDAsInt[3] && sectionsIDAsInt[2] <= sectionsIDAsInt[1] {
			finalValuePartialOverlap += 1
		}

		lowBorder := sectionsIDAsInt[2] - sectionsIDAsInt[0]
		highBorder := sectionsIDAsInt[3] - sectionsIDAsInt[1]
		if isContained(lowBorder, highBorder) {
			finalValueFullOverlap += 1
		}
	}

	return finalValueFullOverlap, finalValuePartialOverlap
}

func Split(r rune) bool {
	return r == ',' || r == '-'
}

func mapStrConv(data []string) ([]int, error) {
	mapped := make([]int, len(data))

	for i, e := range data {
		value, err := strconv.Atoi(e)
		if err != nil {
			return mapped, err
		}
		mapped[i] = value
	}

	return mapped, nil
}

func isContained(lowBorder int, highBorder int) bool {
	// 2-7, 2-8
	if highBorder >= 0 && lowBorder <= 0 {
		return true
	}

	if highBorder <= 0 && lowBorder >= 0 {
		return true
	}

	return false
}
