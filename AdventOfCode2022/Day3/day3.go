package main

import (
    "bufio"
    "fmt"
    "os"
	"unicode"
)

func main() {
	fmt.Println(Challenge_Part2("./chall_input.txt"));
}

func Challenge_Part1(fileName string) int{
	readfile, err := os.Open(fileName);

	if err != nil {
		fmt.Println(err);
		return -1;
	}

	fileScanner := bufio.NewScanner(readfile);
	fileScanner.Split(bufio.ScanLines);
	finalValue := 0;

	for fileScanner.Scan() {
		str := fileScanner.Text();
		histogram := make(map[rune]int);
		for index, ch := range str {
			if v, exists := histogram[ch]; exists {
				if index >= len(str)/2 && v < len(str)/2{
					conversion := convertProperly(ch);
					finalValue += conversion;
					break;
				}
			} else {
				histogram[ch] = index;
			}
		}

	}

	return finalValue
}

func Challenge_Part2(fileName string) int{
	readfile, err := os.Open(fileName);

	if err != nil {
		fmt.Println(err);
		return -1;
	}

	fileScanner := bufio.NewScanner(readfile);
	fileScanner.Split(bufio.ScanLines);
	finalValue := 0;
	histogram := make(map[rune]int);

	for fileScanner.Scan() {
		str := fileScanner.Text();
		foundInString := []rune{};

		for _, ch := range str {
			if contains(foundInString, ch){
				continue;
			}
			foundInString = append(foundInString, ch);
			histogram[ch] += 1;
			if (histogram[ch] == 3){
				finalValue += convertProperly(ch);
				histogram = make(map[rune]int);
				break;
			}
		}

	}

	return finalValue;
}

func convertProperly(ch rune) int {
	if unicode.IsUpper(ch){
		return int(ch) - int('A') + 27;
	}
	return int(ch) - int('a') + 1;
}

func contains(s []rune, e rune) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}