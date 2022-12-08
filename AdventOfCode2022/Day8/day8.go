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

func Challenge(fileName string) (int, int) {
	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()
	treeLine := fileScanner.Text()
	lineIndex := 0
	forest := make([][]int, len(treeLine))

	for true {
		forest[lineIndex] = make([]int, len(treeLine))
		for index, rne := range treeLine {
			value, err := strconv.Atoi(string(rne))
			if err != nil {
				fmt.Printf("%s is not a number in line %d and column %d", string(rne), lineIndex, index)
				return -1, -1
			}
			forest[lineIndex][index] = value
		}

		if !fileScanner.Scan() {
			break
		}

		lineIndex += 1
		treeLine = fileScanner.Text()
	}

	resultChall1 := CountVisibleTrees(forest)
	resultChall2 := ComputeMaxTreeScenicScore(forest)
	return resultChall1, resultChall2
}

func CountVisibleTrees(forest [][]int) int {
	finalCount := 0
	for forestLine := 0; forestLine < len(forest); forestLine++ {
		for treeLine := 0; treeLine < len(forest[forestLine]); treeLine++ {
			if isTreeVisibleLeft(forest, forestLine, treeLine, forest[forestLine][treeLine]) {
				finalCount += 1
				continue
			}
			if isTreeVisibleTop(forest, forestLine, treeLine, forest[forestLine][treeLine]) {
				finalCount += 1
				continue
			}
			if isTreeVisibleRight(forest, forestLine, treeLine, forest[forestLine][treeLine]) {
				finalCount += 1
				continue
			}
			if isTreeVisibleBot(forest, forestLine, treeLine, forest[forestLine][treeLine]) {
				finalCount += 1
				continue
			}
		}
	}
	return finalCount
}

func ComputeMaxTreeScenicScore(forest [][]int) int {
	maxScenicScore := -1
	for forestLine := 0; forestLine < len(forest); forestLine++ {
		for treeLine := 0; treeLine < len(forest[forestLine]); treeLine++ {
			tempScenicScore :=
				computeScenicScoreTop(forest, forestLine, treeLine, forest[forestLine][treeLine], 0) *
					computeScenicScoreRight(forest, forestLine, treeLine, forest[forestLine][treeLine], 0) *
					computeScenicScoreBot(forest, forestLine, treeLine, forest[forestLine][treeLine], 0) *
					computeScenicScoreLeft(forest, forestLine, treeLine, forest[forestLine][treeLine], 0)

			if tempScenicScore > maxScenicScore {
				maxScenicScore = tempScenicScore
			}
		}
	}

	return maxScenicScore
}

func isTreeVisibleTop(forest [][]int, forestLine int, treeLine int, value int) bool {
	if forestLine-1 < 0 {
		return true
	}

	if value <= forest[forestLine-1][treeLine] {
		return false
	}

	return isTreeVisibleTop(forest, forestLine-1, treeLine, value)
}

func isTreeVisibleBot(forest [][]int, forestLine int, treeLine int, value int) bool {
	if forestLine+1 >= len(forest) {
		return true
	}

	if value <= forest[forestLine+1][treeLine] {
		return false
	}

	return isTreeVisibleBot(forest, forestLine+1, treeLine, value)
}

func isTreeVisibleRight(forest [][]int, forestLine int, treeLine int, value int) bool {
	if treeLine+1 >= len(forest[forestLine]) {
		return true
	}

	if value <= forest[forestLine][treeLine+1] {
		return false
	}

	return isTreeVisibleRight(forest, forestLine, treeLine+1, value)
}

func isTreeVisibleLeft(forest [][]int, forestLine int, treeLine int, value int) bool {
	if treeLine-1 < 0 {
		return true
	}

	if value <= forest[forestLine][treeLine-1] {
		return false
	}

	return isTreeVisibleLeft(forest, forestLine, treeLine-1, value)
}

func computeScenicScoreTop(forest [][]int, forestLine int, treeLine int, value int, scenicScore int) int {

	if forestLine-1 < 0 {
		return scenicScore
	}

	scenicScore += 1
	if value <= forest[forestLine-1][treeLine] {
		return scenicScore
	}
	return computeScenicScoreTop(forest, forestLine-1, treeLine, value, scenicScore)
}
func computeScenicScoreBot(forest [][]int, forestLine int, treeLine int, value int, scenicScore int) int {

	if forestLine+1 >= len(forest) {
		return scenicScore
	}

	scenicScore += 1
	if value <= forest[forestLine+1][treeLine] {
		return scenicScore
	}

	return computeScenicScoreBot(forest, forestLine+1, treeLine, value, scenicScore)
}

func computeScenicScoreRight(forest [][]int, forestLine int, treeLine int, value int, scenicScore int) int {

	if treeLine+1 >= len(forest[forestLine]) {
		return scenicScore
	}

	scenicScore += 1
	if value <= forest[forestLine][treeLine+1] {
		return scenicScore
	}

	return computeScenicScoreRight(forest, forestLine, treeLine+1, value, scenicScore)
}

func computeScenicScoreLeft(forest [][]int, forestLine int, treeLine int, value int, scenicScore int) int {

	if treeLine-1 < 0 {
		return scenicScore
	}
	scenicScore += 1
	if value <= forest[forestLine][treeLine-1] {
		return scenicScore
	}

	return computeScenicScoreLeft(forest, forestLine, treeLine-1, value, scenicScore)
}
