package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < n; i++ {
		x := 1 << i
		if x > n {
			break
		}
		fmt.Print(x, " ")
	}
}
