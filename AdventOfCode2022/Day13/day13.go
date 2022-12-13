package main

import (
	"fmt"
	"main/parser"
	"math"
	"os"
	"strings"
)

type State int

const (
	FALSE   State = -1
	NEUTRAL State = 0
	TRUE    State = 1
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
	orderedIndexes := []int{}

	if len(inputs)%2 != 0 {
		fmt.Println("Some pairs are unmatched. Aborting.")
		return -1, -1
	}

	allPairs := []*parser.Sos{}

	addPair2 := parsePairs("[2]")
	addPair6 := parsePairs("[6]")

	for i := 0; i < len(inputs); i++ {
		allPairs = append(allPairs, parsePairs(inputs[i][1:len(inputs[i])-1]))
	}

	allPairs = append(allPairs, addPair2)
	allPairs = append(allPairs, addPair6)

	bubbleSort(allPairs)

	reslt := 1

	for indx, pair := range allPairs {
		if pair == addPair2 {
			reslt *= (indx + 1)
			continue
		}

		if pair == addPair6 {
			reslt *= (indx + 1)
			continue
		}
	}

	for len(inputs) > 0 {

		pair1 := parsePairs(inputs[0][1 : len(inputs[0])-1])
		pair2 := parsePairs(inputs[1][1 : len(inputs[1])-1])
		ordered := comparePairs(pair1, pair2)

		if ordered == TRUE {
			orderedIndexes = append(orderedIndexes, pairIndex)
		}

		if len(inputs) <= 2 {
			break
		}

		inputs = inputs[2:]
		pairIndex += 1
	}

	totalResult := 0
	for _, in := range orderedIndexes {
		totalResult += in
	}

	return totalResult, reslt
}

func bubbleSort(allPairs []*parser.Sos) {
	for i := 0; i < len(allPairs)-1; i++ {
		for j := 0; j < len(allPairs)-i-1; j++ {
			if comparePairs(allPairs[j], allPairs[j+1]) != TRUE {
				temp := allPairs[j]
				allPairs[j] = allPairs[j+1]
				allPairs[j+1] = temp
			}
		}
	}
}

func parsePairs(pair1 string) *parser.Sos {
	sos1 := parser.GetSosFromString(pair1)

	return sos1
}

func comparePairs(pair1 *parser.Sos, pair2 *parser.Sos) State {
	areOrdered := NEUTRAL
	minIndex := int(math.Min(float64(len(pair2.Entries)), float64(len(pair1.Entries))))

	for indx := 0; indx < minIndex; indx++ {
		areOrdered = CompareElems(pair1.Entries[indx], pair2.Entries[indx])
		if areOrdered != NEUTRAL {
			return areOrdered
		}
	}

	if len(pair1.Entries) < len(pair2.Entries) {
		return TRUE
	} else if len(pair1.Entries) == len(pair2.Entries) {
		return NEUTRAL
	}

	return FALSE
}

func CompareElems(elem1 *parser.Element, elem2 *parser.Element) State {
	if elem1.Elemtp == parser.VALUE && elem2.Elemtp == parser.VALUE {
		return giveStateFromComparison(elem1.Value, elem2.Value)
	}

	if elem1.Elemtp == parser.LIST && elem2.Elemtp == parser.LIST {
		tempOrdered := NEUTRAL
		minIndex := int(math.Min(float64(len(elem1.ElemList)), float64(len(elem2.ElemList))))
		for indx := 0; indx < minIndex; indx++ {
			tempOrdered = CompareElems(elem1.ElemList[indx], elem2.ElemList[indx])
			if tempOrdered != NEUTRAL {
				return tempOrdered
			}
		}

		if len(elem1.ElemList) < len(elem2.ElemList) {
			return TRUE
		} else if len(elem1.ElemList) == len(elem2.ElemList) {
			return NEUTRAL
		}

		return FALSE
	}

	if elem1.Elemtp == parser.VALUE {
		elem1.Elemtp = parser.LIST
		newElemValue := &parser.Element{Elemtp: parser.VALUE, Value: elem1.Value}
		elem1.Value = 0
		elem1.ElemList = []*parser.Element{newElemValue}

		return CompareElems(elem1, elem2)
	}

	if elem2.Elemtp == parser.VALUE {
		elem2.Elemtp = parser.LIST
		newElemValue := &parser.Element{Elemtp: parser.VALUE, Value: elem2.Value}
		elem2.Value = 0
		elem2.ElemList = []*parser.Element{newElemValue}

		return CompareElems(elem1, elem2)
	}

	fmt.Println("For some reason, we reached this part ?")
	return FALSE
}

func giveStateFromComparison(value1 int, value2 int) State {
	if value1 < value2 {
		return TRUE
	}

	if value1 > value2 {
		return FALSE
	}

	return NEUTRAL
}

func SplitToBreak(r rune) bool {
	return r == '\n' || r == '\r'
}
