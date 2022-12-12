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

func ChildContains(nodeParent *Node, nodeChild *Node) bool {
	for _, nde := range nodeParent.Children {
		if nde == nodeChild {
			return true
		}
	}

	return false
}
