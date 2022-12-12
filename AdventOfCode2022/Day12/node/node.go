package node

type Coordinate struct {
	XCoord int
	YCoord int
}

type Node struct {
	Name         rune
	Children     []*Node
	PathToAccess int
	Visited      bool
}
