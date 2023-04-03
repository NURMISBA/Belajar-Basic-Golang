package main

import (
	"fmt"
)

func main() {
	fmt.Println("ARITMETIKA PERKALIAN")
	fmt.Println(" ")
	var (
		nilai1 int8 = 4
		nilai2 int8 = 8
	)

	nilai_total := nilai1 * nilai2
	fmt.Println("Hasil dari Perkalian ", nilai1, "*", nilai2, " = ", nilai_total)
}
