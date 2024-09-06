package main

import (
	"fmt"
)

func main() {
	n := vote(0, 0, 1)
	fmt.Println(n)
}
func vote(x int, y int, z int) int {
	if x == y || x == z {
		return x
	} else if y == z {
		return y
	} else {
		return x
	}
}
