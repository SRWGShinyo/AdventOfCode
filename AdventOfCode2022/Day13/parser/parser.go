package parser

import (
	"fmt"
	"strconv"
)

type Entry string
type Symbol rune

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

type Sos struct {
	Entries []*Element
}

func GetSosFromString(sosAsString string) *Sos {

	finalSos := &Sos{}
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

func PrintSos(sosToprint *Sos) {
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
