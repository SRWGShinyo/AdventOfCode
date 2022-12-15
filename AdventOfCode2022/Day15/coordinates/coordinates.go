package coordinates

import (
	"fmt"
	"math"
)

type Coordinate struct {
	XCoord int64
	YCoord int64
}

type XInterval struct {
	XCoordMin int64
	XCoordMax int64
}

func ComputeManhattanDistance(c1 Coordinate, c2 Coordinate) int64 {
	return int64(math.Abs(float64(c1.XCoord)-float64(c2.XCoord)) + math.Abs(float64(c1.YCoord)-float64(c2.YCoord)))
}

func IsComprisedIntoYRange(c1 Coordinate, c2 Coordinate, toCompare Coordinate) bool {
	return c1.YCoord <= toCompare.YCoord && toCompare.YCoord >= c2.YCoord
}

func IsComprisedIntoXRange(c1 Coordinate, c2 Coordinate, toCompare Coordinate) bool {
	return c1.XCoord <= toCompare.XCoord && toCompare.XCoord >= c2.XCoord
}

func Print(it XInterval) {
	fmt.Printf("Min: %d - Max: %d\n", it.XCoordMin, it.XCoordMax)
}

func CompareInterval(i1 XInterval, i2 XInterval) bool {
	if i1.XCoordMin == i2.XCoordMin {
		return i1.XCoordMax < i2.XCoordMax
	}

	return i1.XCoordMin < i2.XCoordMin
}

func CoordinateEquals(c1 Coordinate, c2 Coordinate) bool {
	return c1.XCoord == c2.XCoord && c1.YCoord == c2.YCoord
}

func NumberOfPointsInInterval(i1 XInterval) int64 {
	return i1.XCoordMax - i1.XCoordMin + 1
}

func TryToFuse(i1 XInterval, i2 XInterval) (bool, XInterval) {
	if i2.XCoordMin > i1.XCoordMax {
		return false, i2
	}

	if i2.XCoordMax > i1.XCoordMax {
		return true, XInterval{XCoordMin: i1.XCoordMin, XCoordMax: i2.XCoordMax}
	}

	return true, XInterval{XCoordMin: i1.XCoordMin, XCoordMax: i1.XCoordMax}
}

// This is for sorting operations

type SequenceOrder []XInterval

// By allowing SequenceOrder to implement Len, Less and Swap,
// we can use the sort package to compute the result in O(nlog(n))
func (so SequenceOrder) Len() int           { return len(so) }
func (so SequenceOrder) Less(i, j int) bool { return CompareInterval(so[i], so[j]) }
func (so SequenceOrder) Swap(i, j int)      { so[i], so[j] = so[j], so[i] }
