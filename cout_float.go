package main

import "fmt"

var (
	number float64
)

func init()  {
	fmt.Scan(&number)
}

func main()  {
	if number > 0 {
		number := number * number
		if number > 10000 {
			fmt.Printf("%e", number)
		} else {
			fmt.Printf("%.4f", number)
		}
		//А если число больше чем 10 000 - выводить исходное число в
		// экспоненциальном представлении (знак экспоненты в нижнем регистре).

	} else {
		fmt.Printf("число %.2f не подходит", number)
	}
}