package main

import (
	"fmt"
	"main/parser"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt"))
}

func Challenge(fileName string) (int, int) {
	stringFile, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		return -1, -1
	}

	inputs := strings.FieldsFunc(string(stringFile), SplitToBreak)

	pairIndex := 1
	decoderKey := 1
	numberOfOrderedPairs := 0

	if len(inputs)%2 != 0 {
		fmt.Println("Some pairs are unmatched. Aborting.")
		return -1, -1
	}

	additionalPairs := []*parser.Sequence{parser.ParseSequence("[2]"), parser.ParseSequence("[6]")}
	allPairs := append([]*parser.Sequence{}, additionalPairs...)

	for i := 0; i < len(inputs); i++ {
		allPairs = append(allPairs, parser.ParseSequence(inputs[i][1:len(inputs[i])-1]))
	}

	sort.Sort(parser.SequenceOrder(allPairs))
	for indx, pair := range allPairs {
		if pair == additionalPairs[0] || pair == additionalPairs[1] {
			decoderKey *= (indx + 1)
			continue
		}
	}

	// For part 1
	for len(inputs) > 0 {

		pair1 := parser.ParseSequence(inputs[0][1 : len(inputs[0])-1])
		pair2 := parser.ParseSequence(inputs[1][1 : len(inputs[1])-1])
		ordered := parser.CompareSequence(pair1, pair2)

		if ordered == true {
			numberOfOrderedPairs += pairIndex
		}

		if len(inputs) <= 2 {
			break
		}

		inputs = inputs[2:]
		pairIndex += 1
	}

	return numberOfOrderedPairs, decoderKey
}

func bubbleSort(allPairs []*parser.Sequence) {
	for i := 0; i < len(allPairs)-1; i++ {
		for j := 0; j < len(allPairs)-i-1; j++ {
			if parser.CompareSequence(allPairs[j], allPairs[j+1]) != true {
				temp := allPairs[j]
				allPairs[j] = allPairs[j+1]
				allPairs[j+1] = temp
			}
		}
	}
}

func SplitToBreak(r rune) bool {
	return r == '\n' || r == '\r'
}
