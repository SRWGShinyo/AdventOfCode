package tree

import (
	"fmt"
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
