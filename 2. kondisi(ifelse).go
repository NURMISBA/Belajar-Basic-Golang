package main

import (
	"fmt"
)

func main() {
	fmt.Println("Menentukan nilai terbesar")

	var (
		nilai1, nilai2, nilai3 int
	)

	fmt.Print("Masukkan nilai pertama : ")
	fmt.Scan(&nilai1)

	fmt.Print("Masukkan nilai pertama : ")
	fmt.Scan(&nilai2)

	fmt.Print("Masukkan nilai pertama : ")
	fmt.Scan(&nilai3)

	if nilai1 > nilai2 && nilai1 > nilai3 {
		fmt.Println(nilai1, "Adalah bilangan terbesar ")
	} else if nilai2 > nilai3 && nilai2 > nilai1 {
		fmt.Println(nilai2, "Adalah bilangan terbesar ")
	} else {
		fmt.Println(nilai3, "Adalah Bilangan terbesar")
	}
}
