package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var angka int64

	fmt.Print("Masukkan angka:")
	fmt.Scan(&angka)

	if angka%2 == 0 {
		fmt.Println("ini angka genap")
	} else {
		fmt.Println("ini angga ganjil")
	}
	
}
