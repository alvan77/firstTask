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

	for a, b := range duit {
		fmt.Println(a * b)
	}
}
