package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	numberFirst  string
	numberSecond string
	firstArray   []int
	secondArray  []int
)

func init() {
	fmt.Scan(&numberFirst, &numberSecond)

	firstArray = stringSplitter(numberFirst)
	secondArray = stringSplitter(numberSecond)
}

func main() {
	findMatching(firstArray, secondArray)
}

func stringSplitter(number string) []int {
	var (
		array []int
		split = strings.Split(number, "")
	)

	for i:= 0; i <len(split); i++ {
		element, _ := strconv.Atoi(split[i])
		array = append(array, element)
	}

	return array
}

func findMatching(firstArr []int, secondArr []int) {
	for i:= range firstArr {
		index := indexOf(firstArr[i], secondArr)

		if index != -1 {
			fmt.Printf("%v ",secondArr[index])
		}
	}
}

func indexOf(element int, data []int) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}
