package main

import (
	"bufio"
	"fmt"
	"main/point"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	DOWN  Direction = "D"
	UP    Direction = "U"
	LEFT  Direction = "L"
	RIGHT Direction = "R"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt", 10))
}

func Challenge(fileName string, numberOfNodes int) int {

	if numberOfNodes < 1 {
		fmt.Println("A number of node inferior to 1 renders the whole problem stupid.")
		return -1
	}

	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	individualPoints := []point.Point{point.CreateNewPoint(0, 0)}

	ropeNodesPositions := make([]point.Point, numberOfNodes)
	for fileScanner.Scan() {
		inpt := strings.Split(fileScanner.Text(), " ")
		direction := inpt[0]
		iter, err := strconv.Atoi(inpt[1])
		if err != nil {
			fmt.Printf("%s is not a number, aborting.", inpt[1])
			return -1
		}

		for i := 0; i < iter; i++ {
			ropeNodesPositions[0] = whereToMoveHead(ropeNodesPositions[0], Direction(direction))
			for nodePos := 0; nodePos < len(ropeNodesPositions)-1; nodePos++ {
				if point.AreAdjacent(ropeNodesPositions[nodePos], ropeNodesPositions[nodePos+1]) {
					break
				}
				ropeNodesPositions[nodePos+1] = whereToMoveTail(ropeNodesPositions[nodePos], ropeNodesPositions[nodePos+1])
			}

			if !contains(individualPoints, ropeNodesPositions[len(ropeNodesPositions)-1]) {
				individualPoints = append(individualPoints, ropeNodesPositions[len(ropeNodesPositions)-1])
			}
		}
	}
	return len(individualPoints)
}

func whereToMoveTail(headPosition point.Point, tailPosition point.Point) point.Point {

	if headPosition.Xcoord == tailPosition.Xcoord {
		if headPosition.Ycoord > tailPosition.Ycoord {
			return point.CreateNewPoint(tailPosition.Xcoord, tailPosition.Ycoord+1)
		}
		return point.CreateNewPoint(tailPosition.Xcoord, tailPosition.Ycoord-1)
	}

	if headPosition.Ycoord == tailPosition.Ycoord {
		if headPosition.Xcoord > tailPosition.Xcoord {
			return point.CreateNewPoint(tailPosition.Xcoord+1, tailPosition.Ycoord)
		}
		return point.CreateNewPoint(tailPosition.Xcoord-1, tailPosition.Ycoord)
	}

	diagonalPoint := point.CreateNewPoint(0, 0)
	if headPosition.Xcoord > tailPosition.Xcoord {
		diagonalPoint.Xcoord = 1
	} else {
		diagonalPoint.Xcoord = -1
	}

	if headPosition.Ycoord > tailPosition.Ycoord {
		diagonalPoint.Ycoord = 1
	} else {
		diagonalPoint.Ycoord = -1
	}

	return point.CreateNewPoint(tailPosition.Xcoord+diagonalPoint.Xcoord, tailPosition.Ycoord+diagonalPoint.Ycoord)
}

func whereToMoveHead(headPosition point.Point, whereTo Direction) point.Point {
	switch whereTo {
	case LEFT:
		return point.CreateNewPoint(headPosition.Xcoord-1, headPosition.Ycoord)
	case RIGHT:
		return point.CreateNewPoint(headPosition.Xcoord+1, headPosition.Ycoord)
	case DOWN:
		return point.CreateNewPoint(headPosition.Xcoord, headPosition.Ycoord-1)
	case UP:
		return point.CreateNewPoint(headPosition.Xcoord, headPosition.Ycoord+1)
	}

	fmt.Printf("No direction is given, format %s invalid", whereTo)
	return point.CreateNewPoint(0, 0)
}

func contains(rope []point.Point, newPoint point.Point) bool {
	for _, p := range rope {
		if point.AreEqual(p, newPoint) {
			return true
		}
	}

	return false
}
