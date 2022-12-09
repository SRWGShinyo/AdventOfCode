package point

import (
	"fmt"
	"math"
)

type Point struct {
	Xcoord int
	Ycoord int
}

func AreEqual(p1 Point, p2 Point) bool {
	return (p1.Xcoord == p2.Xcoord) && (p1.Ycoord == p2.Ycoord)
}

func Print(p Point) {
	fmt.Printf("{ x: %d, y: %d }\n", p.Xcoord, p.Ycoord)
}

func CreateNewPoint(x int, y int) Point {
	return Point{Xcoord: x, Ycoord: y}
}

func DistanceBetween(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.Xcoord-p1.Xcoord), 2) + math.Pow(float64(p2.Ycoord-p1.Ycoord), 2))
}

func AreAdjacent(p1 Point, p2 Point) bool {
	return !(DistanceBetween(p1, p2) > math.Sqrt(2))
}
