package main

import (
	"bufio"
	"fmt"
	"main/monkey"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt", 10000))
}

func Challenge(fileName string, roundsNumber int) int {

	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)

	monkeys := []*monkey.Monkey{}
	moduloProduct := big.NewInt(1)

	for fileScanner.Scan() {
		monkInpt := strings.FieldsFunc(fileScanner.Text(), SplitMonkeyInput)
		if monkInpt[0] == "Monkey" {
			monkeyNumber, err := strconv.Atoi(monkInpt[1])
			if err != nil {
				fmt.Printf("%s is not a number, could not create monkey. Aborting\n", monkInpt[1])
				return -1
			}

			fileScanner.Scan()
			monkeyToAdd, erro := CreateMonkeyFromInput(fileScanner)
			if erro == true {
				fmt.Printf("Monkey %d is nil. Aborting.\n", monkeyNumber)
				return -1
			}
			// For part 2 we use a particularity of the modulo
			// for all integers k: (a mod km) mod m = a mod m
			// Finding the LCM between all mod values also works
			moduloProduct.Mul(moduloProduct, big.NewInt(monkeyToAdd.ThrowOperation.Value.Int64()))
			monkeys = append(monkeys, monkeyToAdd)
		}
	}

	fmt.Printf("%d is the moduloProductValue\n", moduloProduct.Int64())

	for i := 0; i < roundsNumber; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkea := monkeys[j]
			for len(monkea.StartingItems) > 0 {
				item := monkea.StartingItems[0]
				monkea.StartingItems = monkea.StartingItems[1:]
				newValueItem := monkey.ApplyMonkeyOp(monkea.StressOperation, item)
				// For exercise two, we apply a particularity of the modulo
				// See https://old.reddit.com/r/adventofcode/comments/zihouc/2022_day_11_part_2_might_need_to_borrow_a_nasa/izrimjo/ for explanation
				newValueItem.Mod(newValueItem, moduloProduct)

				// For exercise 1
				//newValueItem.Div(newValueItem, big.NewInt(3))

				getThrowOpValue := monkey.ApplyMonkeyThrowOp(monkea.ThrowOperation, newValueItem)
				if getThrowOpValue {
					monkeys[monkea.ThrowOperation.MonkeyIfTrue].StartingItems = append(monkeys[monkea.ThrowOperation.MonkeyIfTrue].StartingItems, newValueItem)
				} else {
					monkeys[monkea.ThrowOperation.MonkeyIfFalse].StartingItems = append(monkeys[monkea.ThrowOperation.MonkeyIfFalse].StartingItems, newValueItem)
				}
				monkea.InspectedItemsValue += 1
			}
		}
	}

	//PrintMonkeyMap(monkeys)

	return GetMonkeyBusinessLevel(monkeys)
}

func CreateMonkeyFromInput(fileScanner *bufio.Scanner) (*monkey.Monkey, bool) {
	newMonkey := &monkey.Monkey{}
	inpt := strings.FieldsFunc(fileScanner.Text(), SplitMonkeyInput)
	for len(inpt) > 1 {
		switch inpt[0] {
		case "Starting":
			startingItems, err := GetStartingItemFromStringList(inpt[2:])
			if err {
				fmt.Printf("Could not get starting items. Aborting.\n")
				return newMonkey, true
			}
			newMonkey.StartingItems = append(newMonkey.StartingItems, startingItems...)
		case "Operation":
			stressOp, err := GetStressOpFromInpt(inpt[3:])
			if err {
				fmt.Printf("Could not get stressing operation. Aborting.\n")
				return newMonkey, true
			}
			newMonkey.StressOperation = stressOp
		case "Test":
			throwOp, err := GetThrowOpFromInpt(inpt[1:])
			if err {
				fmt.Printf("Could not get throwing operation. Aborting.\n")
				return newMonkey, true
			}
			newMonkey.ThrowOperation = throwOp
		case "If":
			val, err := strconv.Atoi(inpt[len(inpt)-1])
			if err != nil {
				fmt.Printf("%s is not a number, could not get value for condition %s.Aborting\n", inpt[len(inpt)-1], inpt[1])
				return newMonkey, true
			}

			if inpt[1] == "true" {
				newMonkey.ThrowOperation.MonkeyIfTrue = val
			} else if inpt[1] == "false" {
				newMonkey.ThrowOperation.MonkeyIfFalse = val
			} else {
				fmt.Printf("%s matches neither true nor false. Parsing error ? Aborting.\n", inpt[1])
				return newMonkey, true
			}
		default:
			fmt.Printf("String %s is unmatched. Parsing problem ?", inpt[0])
		}

		fileScanner.Scan()
		inpt = strings.FieldsFunc(fileScanner.Text(), SplitMonkeyInput)
	}

	return newMonkey, false
}

func GetStartingItemFromStringList(itemList []string) ([]*big.Int, bool) {

	startingItems := []*big.Int{}

	for _, value := range itemList {
		item, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("%d is not an int. Parsing error ?\n", item)
		}
		startingItems = append(startingItems, big.NewInt(int64(item)))
	}

	return startingItems, false
}

func GetStressOpFromInpt(inpt []string) (monkey.StressOperationDescriptor, bool) {
	stressOpToReturn := monkey.StressOperationDescriptor{}
	switch inpt[0] {
	case "*":
		stressOpToReturn.Operation = monkey.Operator("MULT")
	case "+":
		stressOpToReturn.Operation = monkey.Operator("ADD")
	default:
		fmt.Printf("Operator %s is not recognized. Aborting.\n", inpt[0])
		return stressOpToReturn, true
	}

	_, err := strconv.Atoi(inpt[1])
	if err != nil && inpt[1] != "old" {
		fmt.Printf("Value %s is not a number nor is it old. Aborting.\n", inpt[1])
		return stressOpToReturn, true
	}

	stressOpToReturn.Value = inpt[1]
	return stressOpToReturn, false
}

func GetThrowOpFromInpt(inpt []string) (monkey.ThrowOperationDescriptor, bool) {
	throwOpToReturn := monkey.ThrowOperationDescriptor{}
	switch inpt[0] {
	case "divisible":
		throwOpToReturn.Operation = monkey.Operator("DIV")
	default:
		fmt.Printf("Operator %s is not recognized. Aborting.\n", inpt[0])
		return throwOpToReturn, true
	}

	value, err := strconv.Atoi(inpt[2])
	if err != nil {
		fmt.Printf("Value %s is not a number. Aborting.\n", inpt[1])
		return throwOpToReturn, true
	}

	throwOpToReturn.Value = big.NewInt(int64(value))
	return throwOpToReturn, false
}

func SplitMonkeyInput(r rune) bool {
	return r == ':' || r == ',' || r == ' ' || r == '='
}

func PrintMonkeyMap(monkeys []*monkey.Monkey) {
	for v, monk := range monkeys {
		fmt.Printf("Printing Monkey %d\n", v)
		monkey.PrintMonkey(*monk)
	}
}

func GetMonkeyBusinessLevel(monkeys []*monkey.Monkey) int {
	monkeyLevel1, monkeyLevel2 := -1, -1
	for m := 0; m < len(monkeys); m++ {
		monk := monkeys[m]
		if monk.InspectedItemsValue > monkeyLevel1 {
			if monkeyLevel1 > monkeyLevel2 {
				monkeyLevel2 = monkeyLevel1
			}
			monkeyLevel1 = monk.InspectedItemsValue
			continue
		}

		if monk.InspectedItemsValue > monkeyLevel2 {
			monkeyLevel2 = monk.InspectedItemsValue
			continue
		}
	}

	return monkeyLevel1 * monkeyLevel2
}
