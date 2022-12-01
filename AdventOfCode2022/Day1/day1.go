package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"sort"
)

func main() {
	fmt.Println(Challenge("./chall1_input.txt"));
}

func Challenge(fileName string) (int, int) {
	readfile, err := os.Open(fileName);

	if err != nil {
		fmt.Println(err);
		return -1, -1;
	}

	fileScanner := bufio.NewScanner(readfile);
	fileScanner.Split(bufio.ScanLines);

	max := -1;
	max3 := []int{0,0,0}
	calories_accumulated := 0;

	for fileScanner.Scan() {
		str := fileScanner.Text();
		calories, err := strconv.Atoi(str);
		if err == nil {
			calories_accumulated += calories;
		} else {
			if calories_accumulated > max {
				max = calories_accumulated;
			}
			if calories_accumulated > max3[0]{
				max3[0] = calories_accumulated;
				sort.Ints(max3);
			}
			calories_accumulated = 0;
		}
	}

	max3Cumulated := 0;

	for _, v := range max3{
		max3Cumulated += v;
	}

	readfile.Close();
	return max, max3Cumulated;
}
