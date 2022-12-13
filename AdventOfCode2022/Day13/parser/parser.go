package parser

import (
	"fmt"
	"math"
	"strconv"
)

type Entry string
type Symbol rune
type State int

const (
	FALSE   State = -1
	NEUTRAL State = 0
	TRUE    State = 1
)

const (
	LIST  Entry = "LIST"
	VALUE Entry = "VALUE"
)

const (
	OPEN_LIST  Symbol = '['
	CLOSE_LIST Symbol = ']'
	COMMA      Symbol = ','
)

type Element struct {
	Elemtp Entry

	ElemList []*Element
	Value    int
}

type Sequence struct {
	Entries []*Element
}

func GetSosFromString(sosAsString string) *Sequence {

	finalSos := &Sequence{}
	tempValue := ""

	for len(sosAsString) > 0 {
		switch sosAsString[0] {
		case byte(OPEN_LIST):
			{
				strVal, elem := createElemListFromString(sosAsString[1:])
				sosAsString = strVal
				finalSos.Entries = append(finalSos.Entries, elem)
			}
		case byte(COMMA):
			{
				newElem := createElemValueFromString(tempValue)
				if newElem != nil {
					finalSos.Entries = append(finalSos.Entries, newElem)
				}

				tempValue = ""
				if len(sosAsString) == 1 {
					sosAsString = ""
					continue
				}
				sosAsString = sosAsString[1:]
			}
		case byte(CLOSE_LIST):
			{
				fmt.Println("This shouldn't happen. Encountered CLOSE_LIST in GetSosFromString")
				if len(sosAsString) == 1 {
					sosAsString = ""
					continue
				}
				sosAsString = sosAsString[1:]
			}
		default:
			{
				tempValue += string(sosAsString[0])
				if len(sosAsString) == 1 {
					newElem := createElemValueFromString(tempValue)
					if newElem != nil {
						finalSos.Entries = append(finalSos.Entries, newElem)
					}
					tempValue = ""
					sosAsString = ""
					continue
				}
				sosAsString = sosAsString[1:]
			}
		}
	}

	return finalSos
}

func createElemListFromString(elemAsList string) (string, *Element) {
	finElem := &Element{Elemtp: LIST}
	tempValue := ""

	for len(elemAsList) > 0 {
		switch elemAsList[0] {
		case byte(OPEN_LIST):
			{
				tempValue = ""
				strVal, elem := createElemListFromString(elemAsList[1:])
				elemAsList = strVal
				finElem.ElemList = append(finElem.ElemList, elem)
			}
		case byte(CLOSE_LIST):
			{
				newElem := createElemValueFromString(tempValue)
				if newElem != nil {
					finElem.ElemList = append(finElem.ElemList, newElem)
				}
				return elemAsList[1:], finElem
			}
		case byte(COMMA):
			{
				newElem := createElemValueFromString(tempValue)
				if newElem != nil {
					finElem.ElemList = append(finElem.ElemList, newElem)
				}
				tempValue = ""
				elemAsList = elemAsList[1:]
			}
		default:
			{
				tempValue += string(elemAsList[0])
				elemAsList = elemAsList[1:]
			}
		}
	}

	return elemAsList, finElem
}

func createElemValueFromString(elemAsValue string) *Element {

	value, err := strconv.Atoi(elemAsValue)
	if err != nil {
		return nil
	}

	finElem := &Element{Elemtp: VALUE, Value: value}
	return finElem
}

func PrintSequence(sosToprint *Sequence) {
	fmt.Printf("[")
	for _, elem := range sosToprint.Entries {
		PrintElem(elem)
	}
	fmt.Printf("]\n")
}

func PrintElem(elemToPrint *Element) {
	fmt.Printf("{ ")
	if elemToPrint.Elemtp == VALUE {
		fmt.Printf("type: %s - value: %d ", elemToPrint.Elemtp, elemToPrint.Value)
	} else {
		fmt.Printf("type: %s - value : [", elemToPrint.Elemtp)
		for _, elem := range elemToPrint.ElemList {
			PrintElem(elem)
		}
		fmt.Printf("]")
	}
	fmt.Printf(" }")
}

func ParseSequence(pair1 string) *Sequence {
	sos1 := GetSosFromString(pair1)

	return sos1
}

func CompareSequence(pair1 *Sequence, pair2 *Sequence) bool {
	areOrdered := NEUTRAL
	minIndex := int(math.Min(float64(len(pair2.Entries)), float64(len(pair1.Entries))))

	for indx := 0; indx < minIndex; indx++ {
		areOrdered = CompareElems(pair1.Entries[indx], pair2.Entries[indx])
		if areOrdered == TRUE {
			return true
		}
		if areOrdered == FALSE {
			return false
		}
	}

	if len(pair1.Entries) < len(pair2.Entries) {
		return true
	}

	return false
}

func CompareElems(elem1 *Element, elem2 *Element) State {
	if elem1.Elemtp == VALUE && elem2.Elemtp == VALUE {
		return giveStateFromComparison(elem1.Value, elem2.Value)
	}

	if elem1.Elemtp == LIST && elem2.Elemtp == LIST {
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

	if elem1.Elemtp == VALUE {
		elem1.Elemtp = LIST
		newElemValue := &Element{Elemtp: VALUE, Value: elem1.Value}
		elem1.Value = 0
		elem1.ElemList = []*Element{newElemValue}

		return CompareElems(elem1, elem2)
	}

	if elem2.Elemtp == VALUE {
		elem2.Elemtp = LIST
		newElemValue := &Element{Elemtp: VALUE, Value: elem2.Value}
		elem2.Value = 0
		elem2.ElemList = []*Element{newElemValue}

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

// This is for sorting operations

type SequenceOrder []*Sequence

// By allowing SequenceOrder to implement Len, Less and Swap,
// we can use the sort package to compute the result in O(nlog(n))
func (so SequenceOrder) Len() int           { return len(so) }
func (so SequenceOrder) Less(i, j int) bool { return CompareSequence(so[i], so[j]) }
func (so SequenceOrder) Swap(i, j int)      { so[i], so[j] = so[j], so[i] }
