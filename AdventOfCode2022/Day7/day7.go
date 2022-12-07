package main

import (
	"bufio"
	"fmt"
	"main/tree"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt", 70000000, 30000000))
}

func Challenge(fileName string, fullStorage int, wantedForUpdate int) int {
	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan() // This is to consume the 'cd /'

	finalTree := ConstructTree(fileScanner)
	tree.PrintTree(finalTree)
	tree.ComputeSize(finalTree)

	tree.GetWantedSize(finalTree, 100000)

	spaceToFree := wantedForUpdate - (fullStorage - finalTree.Size)
	if spaceToFree <= 0 {
		fmt.Println("No need to free up any space.")
		return finalTree.Size
	}

	tree.GetFolderToDelete(finalTree, spaceToFree)

	return finalTree.Size
}

func ConstructTree(fileScanner *bufio.Scanner) *tree.Node {

	finalTree := &tree.Node{Typage: tree.DIR, Parent: nil, Name: "\\", Size: 0, Children: make(map[string]*tree.Node)}
	focusedNode := finalTree
	for fileScanner.Scan() {
		cmd := strings.Split(fileScanner.Text(), " ")
		switch cmd[0] {
		case "$":
			focusedNode = interpretCommand(focusedNode, cmd[1:])
		default:
			focusedNode = interpretLine(focusedNode, cmd)
		}
	}
	return finalTree
}

func interpretCommand(focusedNode *tree.Node, cmd []string) *tree.Node {
	switch cmd[0] {
	case "cd":
		return handleCD(focusedNode, cmd)
	}

	// for ls, we just want to go back up
	return focusedNode
}

func interpretLine(focusedNode *tree.Node, cmd []string) *tree.Node {
	switch cmd[0] {
	case "dir":
		focusedNode = tree.AddNode(focusedNode, cmd[1], 0, tree.DIR)
	default:
		value, err := strconv.Atoi(cmd[0])
		if err == nil {
			focusedNode = tree.AddNode(focusedNode, cmd[1], value, tree.FILE)
		} else {
			fmt.Printf("Error in adding new node file to node %s, value is not a number in %s\n", focusedNode.Name, cmd[0])
		}
	}
	return focusedNode
}

func handleCD(focusedNode *tree.Node, cmd []string) *tree.Node {
	switch cmd[1] {
	case "..":
		return focusedNode.Parent
	default:
		focusedNode = tree.AddNode(focusedNode, cmd[1], 0, tree.DIR)
		return focusedNode.Children[cmd[1]]
	}
}
