// Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
// Входные данные
// 	Программа получает на вход два числа.
// 	Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
// Выходные данные
// 	Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
// 	Цифры выводятся в порядке их нахождения в первом числе.
// 	Внимание! Если вам сложно решить эту задачу, пропустите и изучайте материал дальше, затем вернетесь к этой задаче позже.

// Sample Input:
// 	564 8954
// Sample Output:
// 	5 4

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
	string_splitter(numberFirst, firstArray)
	string_splitter(numberSecond, secondArray)
}

func string_splitter(number string, array []int) {
	split := strings.Split(number, "")
	for i, s := range split {
		array[i], _ = strconv.Atoi(s)
	}

}

func main() {


}
