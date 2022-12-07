package tree

import (
	"fmt"
	"math"
	"strconv"
)

type Filetype int64

const (
	DIR Filetype = iota
	FILE
)

type Node struct {
	Typage   Filetype
	Parent   *Node
	Name     string
	Size     int
	Children map[string]*Node
}

func PrintTree(tree *Node) {
	fmt.Println(FormatNode(tree))
	for _, v := range tree.Children {
		recPrintTree(v, "  ")
	}
}

func recPrintTree(tree *Node, prepend string) {
	fmt.Println(prepend + FormatNode(tree))
	for _, v := range tree.Children {
		recPrintTree(v, prepend+"  ")
	}
}

func FormatNode(node *Node) string {
	if node.Typage == DIR {
		return "- " + node.Name + " (dir)"
	}

	return "- " + node.Name + "(file, size=" + strconv.Itoa(node.Size) + ")"
}

func AddNode(tree *Node, name string, size int, typage Filetype) *Node {
	if _, exists := tree.Children[name]; !exists {
		newNode := &Node{Typage: typage, Parent: tree, Name: name, Size: size, Children: nil}
		if typage == DIR {
			newNode.Children = map[string]*Node{}
		}

		tree.Children[name] = newNode
	}

	return tree
}

func ComputeSize(tree *Node) int {
	for _, v := range tree.Children {
		tree.Size += recComputeSize(v)
	}

	return tree.Size
}

func recComputeSize(tree *Node) int {
	for _, v := range tree.Children {
		tree.Size += recComputeSize(v)
	}

	return tree.Size
}

func GetWantedSize(tree *Node, limit int) int {
	totalSize := 0
	for _, v := range tree.Children {
		totalSize += recGetWantedSize(v, limit, totalSize)
	}

	if tree.Size <= limit && tree.Typage == DIR {
		fmt.Printf("Adding directory %s with size %d\n", tree.Name, tree.Size)
		totalSize += tree.Size
	}

	return totalSize
}

func recGetWantedSize(tree *Node, limit int, totalSize int) int {
	for _, v := range tree.Children {
		recGetWantedSize(v, limit, totalSize)
	}

	if tree.Size <= limit && tree.Typage == DIR {
		fmt.Printf("Adding directory %s with size %d\n", tree.Name, tree.Size)
		return tree.Size + totalSize
	}

	return totalSize
}

func GetFolderToDelete(tree *Node, sizeToDelete int) (string, int) {
	folderName, folderSize := "", math.MaxInt

	for _, v := range tree.Children {
		folderName, folderSize = recGetFolderToDelete(v, sizeToDelete, folderName, folderSize)
	}

	if tree.Typage == DIR && tree.Size >= sizeToDelete && tree.Size < int(folderSize) {
		folderName, folderSize = tree.Name, tree.Size
	}

	fmt.Printf("Finally chose folder %s with size %d\n", folderName, folderSize)
	return folderName, folderSize
}

func recGetFolderToDelete(tree *Node, sizeToDelete int, folderName string, folderSize int) (string, int) {
	for _, v := range tree.Children {
		folderName, folderSize = recGetFolderToDelete(v, sizeToDelete, folderName, folderSize)
	}

	if tree.Typage == DIR && tree.Size >= sizeToDelete && tree.Size < int(folderSize) {
		fmt.Printf("Chose folder %s with size %d\n", folderName, folderSize)
		folderName, folderSize = tree.Name, tree.Size
	}

	return folderName, folderSize
}
