package main

import (
	"bufio"
	"fmt"
	"main/coordinates"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt", 10, 4000000))
}

type SensorAndBeacon struct {
	sensorCoord coordinates.Coordinate
	beaconCoord coordinates.Coordinate
}

type Marker rune

func Challenge(fileName string, lineNumber int64, maxForInterval int64) (int, int64) {
	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	allBeacAndSensInfo := []SensorAndBeacon{}

	for fileScanner.Scan() {
		inpts := FilterStringArrayToIntValue(strings.FieldsFunc(fileScanner.Text(), Split))
		sensorCoordinates := coordinates.Coordinate{XCoord: inpts[0], YCoord: inpts[1]}
		beaconCoordinates := coordinates.Coordinate{XCoord: inpts[2], YCoord: inpts[3]}

		allBeacAndSensInfo = append(allBeacAndSensInfo, SensorAndBeacon{sensorCoord: sensorCoordinates, beaconCoord: beaconCoordinates})
	}

	_, finalNumber := GetIntervalLinesForGivenNumber(allBeacAndSensInfo, lineNumber)
	value := int64(0)

	for i := int64(0); i < maxForInterval; i++ {
		intervalsFin, _ := GetIntervalLinesForGivenNumber(allBeacAndSensInfo, i)
		if len(intervalsFin) > 1 {
			fmt.Printf("x: %d y: %d\n", (intervalsFin[0].XCoordMax+intervalsFin[1].XCoordMin)/2, i)
			value = (intervalsFin[0].XCoordMax+intervalsFin[1].XCoordMin)/2*4000000 + i
			break
		}
	}

	return finalNumber, value
}

func GetIntervalLinesForGivenNumber(sensorsAndBeacon []SensorAndBeacon, lineNumber int64) ([]coordinates.XInterval, int) {
	knownBeacons := []coordinates.Coordinate{}
	intervals := []coordinates.XInterval{}

	for _, sensAndBeacon := range sensorsAndBeacon {
		sensorCoordinates := sensAndBeacon.sensorCoord
		beaconCoordinates := sensAndBeacon.beaconCoord

		manDistance := coordinates.ComputeManhattanDistance(sensorCoordinates, beaconCoordinates)
		if !IsBeaconAlreadyKnow(knownBeacons, beaconCoordinates) {
			knownBeacons = append(knownBeacons, beaconCoordinates)
		}

		// check if the line is in the range of our sensor
		if IsInRange(lineNumber, sensorCoordinates.YCoord-manDistance, sensorCoordinates.YCoord+manDistance) {
			// We know then that all the coordinates from [sensor.X - manDistance, lineNumber]
			// to [sensor.X + manDistance, lineNumber] are in range
			rng := manDistance - int64(math.Abs(float64(sensorCoordinates.YCoord)-float64(lineNumber)))
			intervals = append(intervals, coordinates.XInterval{XCoordMin: sensorCoordinates.XCoord - rng, XCoordMax: sensorCoordinates.XCoord + rng})
		}
	}

	sort.Sort(coordinates.SequenceOrder(intervals))

	finalIntervals := []coordinates.XInterval{}
	targetInterval := intervals[0]

	for i := 1; i < len(intervals); i++ {
		canBeFused, Fusion := coordinates.TryToFuse(targetInterval, intervals[i])
		if canBeFused {
			targetInterval = Fusion
			continue
		}

		finalIntervals = append(finalIntervals, coordinates.XInterval{XCoordMin: targetInterval.XCoordMin, XCoordMax: targetInterval.XCoordMax})
		targetInterval = Fusion
	}

	finalIntervals = append(finalIntervals, targetInterval)
	finalNumber := 0

	for _, interv := range finalIntervals {
		valueToAdd := coordinates.NumberOfPointsInInterval(interv)
		for _, beacon := range knownBeacons {
			if IsInRange(beacon.XCoord, interv.XCoordMin, interv.XCoordMax) && beacon.YCoord == lineNumber {
				valueToAdd -= 1
			}
		}
		finalNumber += int(valueToAdd)
	}

	return finalIntervals, finalNumber
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

func IsInRange(toCompare int64, lowRange int64, highRange int64) bool {
	return lowRange <= toCompare && toCompare <= highRange
}

func IsBeaconAlreadyKnow(knownBeacons []coordinates.Coordinate, newCoord coordinates.Coordinate) bool {
	for _, coord := range knownBeacons {
		if coordinates.CoordinateEquals(coord, newCoord) {
			return true
		}
	}

	return false
}
