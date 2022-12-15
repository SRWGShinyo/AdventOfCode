package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt"))
}

type Marker rune

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

func Split(rne rune) bool {
	return rne == ' ' || rne == '=' || rne == ',' || rne == ':'
}

func FilterStringArrayToIntValue(vs []string) []int64 {
	filtered := make([]int64, 0)
	for _, v := range vs {
		num, err := strconv.Atoi(v)
		if err != nil {
			continue
		}

		filtered = append(filtered, int64(num))
	}

	return filtered
}
