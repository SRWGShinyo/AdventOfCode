package main

import (
	"bufio"
	"fmt"
	"main/node"
	"math"
	"os"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt", 'S', 'E'))
}

func Challenge(fileName string, startNodeValue rune, endNodeValue rune) int {

	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	startCoordinate := node.Coordinate{}
	endCoordinate := node.Coordinate{}
	index := 0
	stringMap := make(map[node.Coordinate]rune)

	for fileScanner.Scan() {
		for indx, rne := range fileScanner.Text() {
			if rne == startNodeValue {
				startCoordinate = node.Coordinate{XCoord: indx, YCoord: index}
				stringMap[node.Coordinate{XCoord: indx, YCoord: index}] = 'a' - 1
			} else if rne == endNodeValue {
				endCoordinate = node.Coordinate{XCoord: indx, YCoord: index}
				stringMap[node.Coordinate{XCoord: indx, YCoord: index}] = 'z' + 1
			} else {
				stringMap[node.Coordinate{XCoord: indx, YCoord: index}] = rne
			}
		}
		index += 1
	}

	startNodes, endNode := constructGraph(stringMap, startCoordinate, endCoordinate)
	minPath := math.MaxInt
	for _, nodeB := range startNodes {
		minPath = int(math.Min(float64(minPath), float64(BFSSearch(nodeB, endNode))))
	}

	return minPath
}

// This is horrible. To optimize later, please.
func constructGraph(stringMap map[node.Coordinate]rune, startCoord node.Coordinate, endCoord node.Coordinate) ([]*node.Node, *node.Node) {
	nodeMap := make(map[node.Coordinate]*node.Node)
	startNode := &node.Node{Name: stringMap[startCoord], Children: []*node.Node{}, PathToAccess: math.MaxInt, Visited: false}
	nodeMap[startCoord] = startNode

	startingPoints := []*node.Node{startNode}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}]; nodeExists {
			if v.Name <= startNode.Name+1 && !node.ChildContains(startNode, v) {
				startNode.Children = append(startNode.Children, v)
				if v.Name == 'a' {
					startingPoints = append(startingPoints, v)
				}
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}], Children: []*node.Node{}, PathToAccess: math.MaxInt, Visited: false}
			if newNode.Name <= startNode.Name+1 && !node.ChildContains(startNode, newNode) {
				startNode.Children = append(startNode.Children, newNode)
			}
			if startNode.Name <= newNode.Name+1 && !node.ChildContains(newNode, startNode) {
				newNode.Children = append(newNode.Children, startNode)
			}

			if newNode.Name == 'a' {
				startingPoints = append(startingPoints, newNode)
			}
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}] = newNode

			startingPoints = recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}, startingPoints)
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}]; nodeExists {
			if v.Name <= startNode.Name+1 && !node.ChildContains(startNode, v) {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}], Children: []*node.Node{}, PathToAccess: math.MaxInt, Visited: false}
			if newNode.Name <= startNode.Name+1 && !node.ChildContains(startNode, newNode) {
				startNode.Children = append(startNode.Children, newNode)
			}
			if startNode.Name <= newNode.Name+1 && !node.ChildContains(newNode, startNode) {
				newNode.Children = append(newNode.Children, startNode)
			}
			if newNode.Name == 'a' {
				startingPoints = append(startingPoints, newNode)
			}
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}] = newNode

			startingPoints = recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}, startingPoints)
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}]; nodeExists {
			if v.Name <= startNode.Name+1 && !node.ChildContains(startNode, v) {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}], Children: []*node.Node{}, PathToAccess: math.MaxInt, Visited: false}
			if newNode.Name <= startNode.Name+1 && !node.ChildContains(startNode, newNode) {
				startNode.Children = append(startNode.Children, newNode)
			}
			if startNode.Name <= newNode.Name+1 && !node.ChildContains(newNode, startNode) {
				newNode.Children = append(newNode.Children, startNode)
			}
			if newNode.Name == 'a' {
				startingPoints = append(startingPoints, newNode)
			}
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}] = newNode

			startingPoints = recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}, startingPoints)
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}]; nodeExists {
			if v.Name <= startNode.Name+1 && !node.ChildContains(startNode, v) {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}], Children: []*node.Node{}, PathToAccess: math.MaxInt, Visited: false}
			if newNode.Name <= startNode.Name+1 && !node.ChildContains(startNode, newNode) {
				startNode.Children = append(startNode.Children, newNode)
			}
			if startNode.Name <= newNode.Name+1 && !node.ChildContains(newNode, startNode) {
				newNode.Children = append(newNode.Children, startNode)
			}
			if newNode.Name == 'a' {
				startingPoints = append(startingPoints, newNode)
			}
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}] = newNode

			startingPoints = recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}, startingPoints)
		}
	}

	return startingPoints, nodeMap[endCoord]
}

func recConstructGraph(stringMap map[node.Coordinate]rune, nodeMap map[node.Coordinate]*node.Node, startNode *node.Node, startCoord node.Coordinate, startingNodes []*node.Node) []*node.Node {
	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}]; nodeExists {
			if v.Name <= startNode.Name+1 && !node.ChildContains(startNode, v) {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}], Children: []*node.Node{}, PathToAccess: math.MaxInt, Visited: false}
			if newNode.Name <= startNode.Name+1 && !node.ChildContains(startNode, newNode) {
				startNode.Children = append(startNode.Children, newNode)
			}
			if startNode.Name <= newNode.Name+1 && !node.ChildContains(newNode, startNode) {
				newNode.Children = append(newNode.Children, startNode)
			}
			if newNode.Name == 'a' {
				startingNodes = append(startingNodes, newNode)
			}
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}] = newNode
			startingNodes = recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}, startingNodes)
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}]; nodeExists {
			if v.Name <= startNode.Name+1 && !node.ChildContains(startNode, v) {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}], Children: []*node.Node{}, PathToAccess: math.MaxInt, Visited: false}
			if newNode.Name <= startNode.Name+1 && !node.ChildContains(startNode, newNode) {
				startNode.Children = append(startNode.Children, newNode)
			}
			if startNode.Name <= newNode.Name+1 && !node.ChildContains(newNode, startNode) {
				newNode.Children = append(newNode.Children, startNode)
			}
			if newNode.Name == 'a' {
				startingNodes = append(startingNodes, newNode)
			}
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}] = newNode

			startingNodes = recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}, startingNodes)
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}]; nodeExists {
			if v.Name <= startNode.Name+1 && !node.ChildContains(startNode, v) {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}], Children: []*node.Node{}, PathToAccess: math.MaxInt, Visited: false}
			if newNode.Name <= startNode.Name+1 && !node.ChildContains(startNode, newNode) {
				startNode.Children = append(startNode.Children, newNode)
			}
			if startNode.Name <= newNode.Name+1 && !node.ChildContains(newNode, startNode) {
				newNode.Children = append(newNode.Children, startNode)
			}
			if newNode.Name == 'a' {
				startingNodes = append(startingNodes, newNode)
			}
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}] = newNode

			startingNodes = recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}, startingNodes)
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}]; nodeExists {
			if v.Name <= startNode.Name+1 && !node.ChildContains(startNode, v) {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}], Children: []*node.Node{}, PathToAccess: math.MaxInt, Visited: false}
			if newNode.Name <= startNode.Name+1 && !node.ChildContains(startNode, newNode) {
				startNode.Children = append(startNode.Children, newNode)
			}
			if startNode.Name <= newNode.Name+1 && !node.ChildContains(newNode, startNode) {
				newNode.Children = append(newNode.Children, startNode)
			}
			if newNode.Name == 'a' {
				startingNodes = append(startingNodes, newNode)
			}
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}] = newNode

			startingNodes = recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}, startingNodes)
		}
	}

	return startingNodes
}

func BFSSearch(startNode *node.Node, endNode *node.Node) int {
	visitedMap := make(map[*node.Node]bool)
	visitedMap[startNode] = true
	deepness := 0
	startNode.PathToAccess = deepness
	nodeQueue := []*node.Node{startNode}
	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		if node == endNode {
			return endNode.PathToAccess
		}
		for _, nodeChild := range node.Children {
			if !visitedMap[nodeChild] {
				nodeChild.PathToAccess = node.PathToAccess + 1
				visitedMap[nodeChild] = true
				nodeQueue = append(nodeQueue, nodeChild)
			}
		}
	}

	fmt.Println("End node wans't found...")
	return math.MaxInt
}
