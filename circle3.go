package main

import (
	"fmt"
)

func main() {
	var (
		acc    int = 0
		index  int
		target int
	)

	fmt.Scan(&index)

	for i := 1; i <= index; i++ {
		fmt.Scan(&target)

		if target >= 10 && target <= 100 && target%8 == 0 {
			acc += target
		}
	}

	fmt.Println(acc)
}
