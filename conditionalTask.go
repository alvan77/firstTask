package main

import "fmt"

func main() {
	var nilai int

	fmt.Println("Masukkan nilai yg anda inginkan :")
	fmt.Scan(&nilai)
	if nilai <= 40 {
		fmt.Printf("nilai anda adalah %d, berusaha lagi\n", nilai)
	} else if nilai <= 70 {
		fmt.Printf("nilai anda adalah %d, sudah cukup\n", nilai)
	} else {
		fmt.Printf("Nilai anda adlah %d, anda luar biasa\n", nilai)
	}
}
