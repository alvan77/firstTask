package main

import (
	"fmt"
)

func main() {
	var money = []map[string]int{
		{"Rupiah": 100000, "Jumlah": 8},
		{"Dollar": 100, "Jumlah": 5},
	}

	var a, b, c, d int
	a = money[0]["Rupiah"]
	b = money[0]["Jumlah"]
	c = money[1]["Dollar"]
	d = money[1]["Jumlah"]
	fmt.Println(a * b)
	fmt.Println(c * d)
}
