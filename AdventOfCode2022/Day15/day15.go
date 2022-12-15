package main

import (
	"bufio"
	"fmt"
	"main/coordinates"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Challenge("./example_input.txt"))
}

type Marker rune

const (
	EMPTY   Marker = '.'
	SCANNED Marker = '#'
	SENSOR  Marker = 'S'
	BEACON  Marker = 'B'
)

func Challenge(fileName string) int {
	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	scanningMap := make(map[int64]map[int64]rune)
	highestCoordinates := coordinates.Coordinate{XCoord: math.MinInt, YCoord: math.MaxInt}
	lowestCoordinates := coordinates.Coordinate{XCoord: math.MaxInt, YCoord: math.MaxInt}

	for fileScanner.Scan() {
		inpts := FilterStringArrayToIntValue(strings.FieldsFunc(fileScanner.Text(), Split))
		sensorCoordinates := coordinates.Coordinate{XCoord: inpts[0], YCoord: inpts[1]}
		beaconCoordinates := coordinates.Coordinate{XCoord: inpts[2], YCoord: inpts[3]}

		scannerMap, tempHighAndLow := AddSensorRangeAndBeacon(scanningMap, sensorCoordinates, beaconCoordinates)

		scanningMap = scannerMap
		highestCoordinates = coordinates.GetHighestFromBoth(highestCoordinates, tempHighAndLow[0])
		lowestCoordinates = coordinates.GetLowestFromBoth(lowestCoordinates, tempHighAndLow[1])

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
			fmt.Printf("%s is not an int, continuing.\n", v)
			continue
		}

		filtered = append(filtered, int64(num))
	}

	return filtered
}

func DoesEntryExistInScannerMapX(maps map[int64]map[int64]rune, XCoord int64) bool {
	if _, exists := maps[XCoord]; exists {
		return true
	}

	return false
}

func DoesEntryExistInScannerMapY(maps map[int64]map[int64]rune, XCoord int64, YCoord int64) bool {
	if _, exists := maps[XCoord][YCoord]; exists {
		return true
	}

	return false
}

func AddSensorRangeAndBeacon(mapScanner map[int64]map[int64]rune, sensorCoordinates coordinates.Coordinate, beaconCoordinate coordinates.Coordinate) (map[int64]map[int64]rune, []coordinates.Coordinate) {
	tempHighAndLow := []coordinates.Coordinate{{XCoord: math.MinInt, YCoord: math.MinInt}, {XCoord: math.MaxInt, YCoord: math.MinInt}}

	return mapScanner, tempHighAndLow
}

func AddCoordinateToMap(mapScanner map[int64]map[int64]rune, coordinateToAdd coordinates.Coordinate, symbolToAdd Marker) map[int64]map[int64]rune {
	if !DoesEntryExistInScannerMapX(mapScanner, coordinateToAdd.XCoord) {
		mapScanner[coordinateToAdd.XCoord] = make(map[int64]rune)
		mapScanner[coordinateToAdd.XCoord][coordinateToAdd.YCoord] = rune(symbolToAdd)

		return mapScanner
	}

	if !DoesEntryExistInScannerMapY(mapScanner, coordinateToAdd.XCoord, coordinateToAdd.YCoord) {
		mapScanner[coordinateToAdd.XCoord][coordinateToAdd.YCoord] = rune(symbolToAdd)
		return mapScanner
	}

	actualValue := mapScanner[coordinateToAdd.XCoord][coordinateToAdd.YCoord]
	if actualValue == rune(BEACON) || actualValue == rune(SENSOR) {
		fmt.Printf("There is already a %s at position %d, %d. Not adding %s.\n", string(actualValue), coordinateToAdd.XCoord, coordinateToAdd.YCoord, string(symbolToAdd))
		return mapScanner
	}

	if actualValue == rune(SCANNED) && symbolToAdd == EMPTY {
		fmt.Printf("There is already a %s at position %d, %d. Not adding %s.\n", string(actualValue), coordinateToAdd.XCoord, coordinateToAdd.YCoord, string(symbolToAdd))
		return mapScanner
	}

	mapScanner[coordinateToAdd.XCoord][coordinateToAdd.YCoord] = rune(symbolToAdd)
	return mapScanner
}
