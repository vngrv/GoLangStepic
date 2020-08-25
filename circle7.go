// Вклад в банке составляет x рублей.
// Ежегодно он увеличивается на p процентов,
// после чего дробная часть копеек отбрасывается.
// Каждый год сумма вклада становится больше.
// Определите, через сколько лет вклад составит не менее y рублей.
// Входные данные
// 	-	Программа получает на вход три натуральных числа: x, p, y.
// Выходные данные
// 	-	Программа должна вывести одно целое число.

// Sample Input:
// 	100 - вклад
// 	10 	- процент увелечения
// 	200 - цель
// Sample Output:
// 	8 	- втф?

package main

import "fmt"

var (
	deposit int
	persent int
	target  int
	counter int
)

func init() {
	fmt.Scan(&deposit, &persent, &target)
}

func main() {
	for ; deposit <= target; deposit += deposit * persent / 100 {
		counter++
	}

	fmt.Println(counter)
}
