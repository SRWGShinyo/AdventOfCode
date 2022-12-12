package main

import (
	"bufio"
	"fmt"
	"main/node"
	"os"
)

func main() {
	fmt.Println(Challenge("./example_input.txt", 'S', 'E'))
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
	index := 0
	stringMap := make(map[node.Coordinate]rune)

	for fileScanner.Scan() {
		for indx, rne := range fileScanner.Text() {
			if rne == startNodeValue {
				startCoordinate = node.Coordinate{XCoord: indx, YCoord: index}
				stringMap[node.Coordinate{XCoord: indx, YCoord: index}] = 'a' - 1
			} else if rne == endNodeValue {
				stringMap[node.Coordinate{XCoord: indx, YCoord: index}] = 'z' + 1
			} else {
				stringMap[node.Coordinate{XCoord: indx, YCoord: index}] = rne
			}
		}
		index += 1
	}

	startNode := constructGraph(stringMap, startCoordinate)
	for _, node := range startNode.Children {
		fmt.Println(string(node.Name))
	}

	return updateShortestPath(startNode, endNodeValue)
}

func constructGraph(stringMap map[node.Coordinate]rune, startCoord node.Coordinate) *node.Node {
	nodeMap := make(map[node.Coordinate]*node.Node)
	startNode := &node.Node{Name: stringMap[startCoord], Children: []*node.Node{}, PathToAccess: 0, Visited: false}
	nodeMap[startCoord] = startNode

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}]; nodeExists {
			if v.Name <= startNode.Name+1 {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}], Children: []*node.Node{}, PathToAccess: 0, Visited: false}
			newNode.Children = append(newNode.Children, startNode)
			startNode.Children = append(startNode.Children, newNode)
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}] = newNode

			recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord})
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}]; nodeExists {
			if v.Name <= startNode.Name+1 {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}], Children: []*node.Node{}, PathToAccess: 0, Visited: false}
			newNode.Children = append(newNode.Children, startNode)
			startNode.Children = append(startNode.Children, newNode)
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}] = newNode

			recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord})
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}]; nodeExists {
			if v.Name <= startNode.Name+1 {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}], Children: []*node.Node{}, PathToAccess: 0, Visited: false}
			newNode.Children = append(newNode.Children, startNode)
			startNode.Children = append(startNode.Children, newNode)
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}] = newNode

			recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1})
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}]; nodeExists {
			if v.Name <= startNode.Name+1 {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}], Children: []*node.Node{}, PathToAccess: 0, Visited: false}
			newNode.Children = append(newNode.Children, startNode)
			startNode.Children = append(startNode.Children, newNode)
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}] = newNode

			recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1})
		}
	}

	return nodeMap[startCoord]
}

func recConstructGraph(stringMap map[node.Coordinate]rune, nodeMap map[node.Coordinate]*node.Node, startNode *node.Node, startCoord node.Coordinate) {
	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}]; nodeExists {
			if v.Name <= startNode.Name+1 {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}], Children: []*node.Node{}, PathToAccess: 0, Visited: false}
			newNode.Children = append(newNode.Children, startNode)
			startNode.Children = append(startNode.Children, newNode)
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord}] = newNode
			recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord + 1, YCoord: startCoord.YCoord})
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}]; nodeExists {
			if v.Name <= startNode.Name+1 {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}], Children: []*node.Node{}, PathToAccess: 0, Visited: false}
			newNode.Children = append(newNode.Children, startNode)
			startNode.Children = append(startNode.Children, newNode)
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord}] = newNode

			recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord - 1, YCoord: startCoord.YCoord})
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}]; nodeExists {
			if v.Name <= startNode.Name+1 {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}], Children: []*node.Node{}, PathToAccess: 0, Visited: false}
			newNode.Children = append(newNode.Children, startNode)
			startNode.Children = append(startNode.Children, newNode)
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1}] = newNode

			recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord + 1})
		}
	}

	if _, exists := stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}]; exists {
		if v, nodeExists := nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}]; nodeExists {
			if v.Name <= startNode.Name+1 {
				startNode.Children = append(startNode.Children, v)
			}
		} else {
			newNode := &node.Node{Name: stringMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}], Children: []*node.Node{}, PathToAccess: 0, Visited: false}
			newNode.Children = append(newNode.Children, startNode)
			startNode.Children = append(startNode.Children, newNode)
			nodeMap[node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1}] = newNode

			recConstructGraph(stringMap, nodeMap, newNode, node.Coordinate{XCoord: startCoord.XCoord, YCoord: startCoord.YCoord - 1})
		}
	}
}

func updateShortestPath(startNode *node.Node, endValue rune) int {
	return -1
}
