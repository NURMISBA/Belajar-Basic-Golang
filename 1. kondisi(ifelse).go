package main

import (
	"fmt"
)

func main() {
	fmt.Println(" ")
	fmt.Println(" Mencari Bilangan Genap dan Ganjil Menggunakan Kondisi")
	var (
		x int
	)
	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println(x, " Merupakan Bilangan Genap")
	} else {
		fmt.Println(x, "Merupakan Bilangan Ganjil")
	}
}
