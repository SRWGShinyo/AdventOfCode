package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

	cavernSlice := make(map[int]map[int]string)

	//set up start point
	cavernSlice[500] = make(map[int]string)
	cavernSlice[500][0] = "+"

	lowestXForARock, highestXForARock := math.MaxInt, math.MinInt
	lowestYForARock, highestYForARock := 0, math.MinInt

	newCoords := Coordinate{XCoord: -1, YCoord: -1}

	for fileScanner.Scan() {
		inpts := strings.FieldsFunc(fileScanner.Text(), Split)
		pastCoords := Coordinate{XCoord: -1, YCoord: -1}
		for _, inpt := range inpts {
			newCoords = RetrieveCoordinates(strings.Split(inpt, ","))

			if newCoords.XCoord == -1 || newCoords.YCoord == -1 {
				return -1
			}

			highestXForARock, lowestXForARock = UpdateHighestAndLowestRockXCoord(newCoords.XCoord, highestXForARock, lowestXForARock)
			highestYForARock, lowestYForARock = UpdateHighestAndLowestRockYCoord(newCoords.YCoord, highestYForARock, lowestYForARock)

			if pastCoords.XCoord == -1 {
				cavernSlice = UpdateCaveMap(cavernSlice, newCoords, newCoords)
			} else {
				cavernSlice = UpdateCaveMap(cavernSlice, newCoords, pastCoords)
			}

			pastCoords = Coordinate{XCoord: newCoords.XCoord, YCoord: newCoords.YCoord}
		}
	}

	lowestCoords := Coordinate{XCoord: lowestXForARock - 5000, YCoord: 0}
	highestCoords := Coordinate{XCoord: highestXForARock + 5000, YCoord: highestYForARock + 1}

	cavernSlice = FullCavernSliceMap(cavernSlice, lowestCoords, highestCoords)

	numberOfGrain := 1
	startPoint := Coordinate{XCoord: 500, YCoord: 0}
	for true {

		newSlice, hasTouchedEntry := MakeGrainFallChall2(cavernSlice, startPoint, lowestCoords, highestCoords)
		if hasTouchedEntry {
			break
		}

		cavernSlice = newSlice
		numberOfGrain += 1
	}

	//PrintCavernSlice(cavernSlice, lowestCoords, highestCoords)

	return numberOfGrain
}

func Split(rne rune) bool {
	return rne == ' ' || rne == '-' || rne == '>'
}

func MakeGrainFallChall2(cavernSlice map[int]map[int]string, startCoord Coordinate, lowestCoords Coordinate, highestCoords Coordinate) (map[int]map[int]string, bool) {
	grainCoord := Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord}

	for true {
		lowerCheck, newGrainCoord := checkLowerChall2(cavernSlice, grainCoord, lowestCoords, highestCoords)
		if lowerCheck {
			grainCoord = newGrainCoord
			continue
		}

		diagLeftCheck, newGrainCoord := checkDiagLeftChall2(cavernSlice, grainCoord, lowestCoords, highestCoords)
		if diagLeftCheck {
			grainCoord = newGrainCoord
			continue
		}

		diagRightCheck, newGrainCoord := checkDiagRightChall2(cavernSlice, grainCoord, lowestCoords, highestCoords)
		if diagRightCheck {
			grainCoord = newGrainCoord
			continue
		}

		cavernSlice[grainCoord.XCoord][grainCoord.YCoord] = "o"
		if grainCoord.XCoord == startCoord.XCoord && grainCoord.YCoord == startCoord.YCoord {
			return cavernSlice, true
		}

		return cavernSlice, false
	}

	return cavernSlice, false
}
func checkLowerChall2(cavernSlice map[int]map[int]string, grainCoord Coordinate, lowestCoords Coordinate, highestCoorsd Coordinate) (bool, Coordinate) {
	grainCoord.YCoord += 1
	if grainCoord.YCoord >= highestCoorsd.YCoord+2 {
		return false, grainCoord
	}

	return cavernSlice[grainCoord.XCoord][grainCoord.YCoord] == ".", grainCoord
}

func checkDiagLeftChall2(cavernSlice map[int]map[int]string, grainCoord Coordinate, lowestCoords Coordinate, highestCoorsd Coordinate) (bool, Coordinate) {

	grainCoord.YCoord += 1
	grainCoord.XCoord -= 1

	if grainCoord.YCoord >= highestCoorsd.YCoord+2 {
		return false, grainCoord
	}

	return cavernSlice[grainCoord.XCoord][grainCoord.YCoord] == ".", grainCoord
}

func checkDiagRightChall2(cavernSlice map[int]map[int]string, grainCoord Coordinate, lowestCoords Coordinate, highestCoorsd Coordinate) (bool, Coordinate) {
	grainCoord.YCoord += 1
	grainCoord.XCoord += 1

	if grainCoord.YCoord >= highestCoorsd.YCoord+2 {
		return false, grainCoord
	}
	return cavernSlice[grainCoord.XCoord][grainCoord.YCoord] == ".", grainCoord
}

func MakeGrainFallChall1(cavernSlice map[int]map[int]string, startCoord Coordinate, lowestCoords Coordinate, highestCoords Coordinate) (map[int]map[int]string, bool) {
	grainCoord := Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord}
	isRested := false

	for !isRested {
		lowerCheck, newGrainCoord := checkLowerChall1(cavernSlice, grainCoord, lowestCoords, highestCoords)
		if lowerCheck {
			if newGrainCoord.XCoord == -1 {
				return cavernSlice, true
			}

			grainCoord = newGrainCoord
			continue
		}

		diagLeftCheck, newGrainCoord := checkDiagLeftChall1(cavernSlice, grainCoord, lowestCoords, highestCoords)
		if diagLeftCheck {
			if newGrainCoord.XCoord == -1 {
				return cavernSlice, true
			}

			grainCoord = newGrainCoord
			continue
		}

		diagRightCheck, newGrainCoord := checkDiagRightChall1(cavernSlice, grainCoord, lowestCoords, highestCoords)
		if diagRightCheck {
			if newGrainCoord.XCoord == -1 {
				return cavernSlice, true
			}

			grainCoord = newGrainCoord
			continue
		}

		cavernSlice[grainCoord.XCoord][grainCoord.YCoord] = "o"
		return cavernSlice, false
	}

	return cavernSlice, false
}
func checkLowerChall1(cavernSlice map[int]map[int]string, grainCoord Coordinate, lowestCoords Coordinate, highestCoorsd Coordinate) (bool, Coordinate) {
	grainCoord.YCoord += 1
	if checkOutOfBound(grainCoord, lowestCoords, highestCoorsd) {
		return true, Coordinate{XCoord: -1, YCoord: -1}
	}
	return cavernSlice[grainCoord.XCoord][grainCoord.YCoord] == ".", grainCoord
}

func checkDiagLeftChall1(cavernSlice map[int]map[int]string, grainCoord Coordinate, lowestCoords Coordinate, highestCoorsd Coordinate) (bool, Coordinate) {

	grainCoord.YCoord += 1
	grainCoord.XCoord -= 1

	if checkOutOfBound(grainCoord, lowestCoords, highestCoorsd) {
		return true, Coordinate{XCoord: -1, YCoord: -1}
	}

	return cavernSlice[grainCoord.XCoord][grainCoord.YCoord] == ".", grainCoord
}

func checkDiagRightChall1(cavernSlice map[int]map[int]string, grainCoord Coordinate, lowestCoords Coordinate, highestCoorsd Coordinate) (bool, Coordinate) {
	grainCoord.YCoord += 1
	grainCoord.XCoord += 1

	if checkOutOfBound(grainCoord, lowestCoords, highestCoorsd) {
		return true, Coordinate{XCoord: -1, YCoord: -1}
	}
	return cavernSlice[grainCoord.XCoord][grainCoord.YCoord] == ".", grainCoord
}

func checkOutOfBound(grainCoord Coordinate, lowestCoords Coordinate, highestCoords Coordinate) bool {
	if grainCoord.XCoord > highestCoords.XCoord || grainCoord.XCoord < lowestCoords.XCoord {
		return true
	}

	if grainCoord.YCoord > highestCoords.YCoord || grainCoord.YCoord < lowestCoords.YCoord {
		return true
	}
	return false
}

func FullCavernSliceMap(cavernSlice map[int]map[int]string, lowestCoordinates Coordinate, highestCoordinates Coordinate) map[int]map[int]string {
	for yc := lowestCoordinates.YCoord; yc < highestCoordinates.YCoord+1; yc++ {
		for xc := lowestCoordinates.XCoord; xc < highestCoordinates.XCoord+1; xc++ {
			if _, exists := cavernSlice[xc]; !exists {
				cavernSlice[xc] = make(map[int]string)
			}
			if _, exists := cavernSlice[xc][yc]; !exists {
				cavernSlice[xc][yc] = "."
			}
		}
	}
	return cavernSlice
}

func RetrieveCoordinates(coords []string) Coordinate {
	XCoord, err := strconv.Atoi(coords[0])
	if err != nil {
		fmt.Printf("%s is not a number. Aborting.\n", coords[0])
		return Coordinate{XCoord: -1, YCoord: -1}
	}

	YCoord, err := strconv.Atoi(coords[1])
	if err != nil {
		fmt.Printf("%s is not a number. Aborting.\n", coords[1])
		return Coordinate{XCoord: -1, YCoord: -1}
	}

	return Coordinate{XCoord: XCoord, YCoord: YCoord}
}

func UpdateHighestAndLowestRockXCoord(coordX int, actualHighest int, actualLowest int) (int, int) {
	if coordX < actualLowest {
		actualLowest = coordX
	}

	if coordX > actualHighest {
		actualHighest = coordX
	}

	return actualHighest, actualLowest
}

func UpdateHighestAndLowestRockYCoord(coordY int, actualHighest int, actualLowest int) (int, int) {
	if coordY < actualLowest {
		actualLowest = coordY
	}

	if coordY > actualHighest {
		actualHighest = coordY
	}

	return actualHighest, actualLowest
}

func UpdateCaveMap(cavernSlice map[int]map[int]string, newCoord Coordinate, pastCoord Coordinate) map[int]map[int]string {

	if pastCoord.XCoord > newCoord.XCoord {
		for xc := pastCoord.XCoord; xc > newCoord.XCoord-1; xc-- {
			if _, exists := cavernSlice[xc]; !exists {
				cavernSlice[xc] = make(map[int]string)
			}

			cavernSlice[xc][newCoord.YCoord] = "#"
		}
	} else {
		for xc := pastCoord.XCoord; xc < newCoord.XCoord+1; xc++ {
			if _, exists := cavernSlice[xc]; !exists {
				cavernSlice[xc] = make(map[int]string)
			}

			cavernSlice[xc][newCoord.YCoord] = "#"
		}
	}

	if pastCoord.YCoord > newCoord.YCoord {
		for yc := pastCoord.YCoord; yc > newCoord.YCoord-1; yc-- {
			if _, exists := cavernSlice[newCoord.XCoord]; !exists {
				cavernSlice[newCoord.XCoord] = make(map[int]string)
			}
			cavernSlice[newCoord.XCoord][yc] = "#"
		}
	} else {
		for yc := pastCoord.YCoord; yc < newCoord.YCoord+1; yc++ {
			if _, exists := cavernSlice[newCoord.XCoord]; !exists {
				cavernSlice[newCoord.XCoord] = make(map[int]string)
			}

			cavernSlice[newCoord.XCoord][yc] = "#"
		}
	}

	return cavernSlice
}

func PrintCavernSlice(cavernSlice map[int]map[int]string, lowestCoordinates Coordinate, highestCoordinates Coordinate) {
	fmt.Printf("%d, %d, %d,%d \n", lowestCoordinates.XCoord, lowestCoordinates.YCoord, highestCoordinates.XCoord, highestCoordinates.YCoord)
	for yc := lowestCoordinates.YCoord; yc < highestCoordinates.YCoord+1; yc++ {
		for xc := lowestCoordinates.XCoord; xc < highestCoordinates.XCoord+1; xc++ {
			fmt.Printf(cavernSlice[xc][yc])
		}
		fmt.Println()
	}
}
