package coordinates

type Coordinate struct {
	XCoord int64
	YCoord int64
}

func Max(x int64, y int64) int64 {
	if x > y {
		return x
	}

	return y
}

func Min(x int64, y int64) int64 {
	if x < y {
		return x
	}

	return y
}

func GetHighestFromBoth(c1 Coordinate, c2 Coordinate) Coordinate {
	return Coordinate{XCoord: Max(c1.XCoord, c2.XCoord), YCoord: Max(c1.YCoord, c2.YCoord)}
}

func GetLowestFromBoth(c1 Coordinate, c2 Coordinate) Coordinate {
	return Coordinate{XCoord: Min(c1.XCoord, c2.XCoord), YCoord: Min(c1.YCoord, c2.YCoord)}
}
