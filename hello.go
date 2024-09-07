package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumInt(1, 0))

	print(5, 6)
}
func sumInt(a ...int) (n int, sum int) {
	for _, v := range a {
		n = n + 1
		sum = sum + v
	}
	return n, sum
}
