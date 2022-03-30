package main

import (
	"fmt"
)

func main() {

	duit := map[int]int{
		1000:  2,
		5000:  4,
		10000: 7,
	}
	var total int
	for a, b := range duit {
		total += a * b
		fmt.Println(a * b)
	}
	fmt.Println(total)
}
