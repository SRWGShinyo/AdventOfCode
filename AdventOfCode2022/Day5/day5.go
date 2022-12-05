package main

import (
	"bufio"
	"fmt"
	"main/container"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(Challenge("./chall_input.txt"))

}

func Challenge(fileName string) string {
	readfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}

	fileScanner := bufio.NewScanner(readfile)
	fileScanner.Split(bufio.ScanLines)
	fileScanner.Scan()

	str := fileScanner.Text()
	stockPiles := []string{}
	pileNumber := 0

	for len(str) != 0 {
		// just before the last
		if !unicode.IsDigit(rune(str[len(str)-2])) {
			stockPiles = append(stockPiles, str)
		} else {
			pileNumber, err = strconv.Atoi(string(str[len(str)-2]))
			if err != nil {
				fmt.Println(err)
				return "ERROR"
			}
		}
		fileScanner.Scan()
		str = fileScanner.Text()
	}

	properStacks := parseInput(stockPiles, pileNumber)

	for fileScanner.Scan() {
		str := fileScanner.Text()
		if len(str) == 0 {
			break
		}
		number, from, to := getNumbersFromInstruction(str)
		if number == -1 {
			fmt.Println("Error in instruction parsing")
			return "ERROR"
		}

		properStacks = CrateMover9001(number, from, to, properStacks)
		if properStacks == nil {
			fmt.Println("An Error occured while moving the crates")
			return "ERROR"
		}
	}

	return getTopStacks(properStacks)
}

func CrateMover9000(number int, from int, to int, properStacks []container.Stack[string]) []container.Stack[string] {
	for i := 0; i < number; i++ {
		crate, ableToPop := properStacks[from-1].Pop()
		if !ableToPop {
			fmt.Printf("Stack %d is empty, impossible to move crates.\n", from)
			break
		}
		properStacks[to-1].Push(crate)
	}

	return properStacks
}

func CrateMover9001(number int, from int, to int, properStacks []container.Stack[string]) []container.Stack[string] {
	crate, ableToPop := properStacks[from-1].PopMultiple(number)
	if !ableToPop {
		fmt.Printf("Stack %d does not contain enough crates, impossible to move.\n", from)
		return nil
	}

	for ind := len(crate) - 1; ind >= 0; ind-- {
		properStacks[to-1].Push(crate[ind])
	}

	// using pushMultiple would result in the crate at position 'from -1' to become
	// a copy of the crate at position 'to-1' before insertion...investigate
	//properStacks[to-1].PushMultiple(crate)

	return properStacks
}

func parseInput(stockPiles []string, number int) []container.Stack[string] {
	propserStacks := make([]container.Stack[string], number)

	for s := len(stockPiles) - 1; s >= 0; s-- {
		for i := 0; i < number; i++ {
			if unicode.IsLetter(rune(stockPiles[s][1+i*4])) {
				propserStacks[i].Push(string(stockPiles[s][1+i*4]))
			}
		}
	}

	return propserStacks
}

func getNumbersFromInstruction(instruct string) (int, int, int) {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(instruct, -1)
	numberValue, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Println(instruct)
		fmt.Println(err)
		return -1, -1, -1
	}

	fromValue, err := strconv.Atoi(numbers[1])
	if err != nil {
		fmt.Println(instruct)
		fmt.Println(err)
		return -1, -1, -1
	}

	toValue, err := strconv.Atoi(numbers[2])
	if err != nil {
		fmt.Println(instruct)
		fmt.Println(err)
		return -1, -1, -1
	}

	return numberValue, fromValue, toValue
}

func getTopStacks(piles []container.Stack[string]) string {

	result := ""

	for _, pile := range piles {
		result += pile.Peek()
	}

	return result
}

func printAllStacks(properStacks []container.Stack[string]) {
	for i := 0; i < len(properStacks); i++ {
		fmt.Printf("print stack %d\n", i)
		fmt.Println(properStacks[i])
	}
}
