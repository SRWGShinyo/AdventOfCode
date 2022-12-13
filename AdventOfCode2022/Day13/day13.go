package main

import (
	"fmt"
	"main/parser"
	"os"
	"strings"
)

func main() {
	fmt.Println(Challenge("./example_input.txt"))
}

func Challenge(fileName string) int {
	stringFile, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	inputs := strings.FieldsFunc(string(stringFile), SplitToBreak)
	fmt.Println(inputs)

	if len(inputs)%2 != 0 {
		fmt.Println("Some pairs are unmatched. Aborting.")
		return -1
	}

	for len(inputs) > 0 {

		parsePairs(inputs[0][1:len(inputs[0])-1], inputs[1][1:len(inputs[1])-1])
		if len(inputs) <= 2 {
			break
		}

		inputs = inputs[2:]
	}

	return -1
}

func parsePairs(pair1 string, pair2 string) {
	sos1 := parser.GetSosFromString(pair1)
	sos2 := parser.GetSosFromString(pair2)

	parser.PrintSos(sos1)
	parser.PrintSos(sos2)
}

func SplitToBreak(r rune) bool {
	return r == '\n' || r == '\r'
}
