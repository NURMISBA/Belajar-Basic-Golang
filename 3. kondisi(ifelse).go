package main

import (
	"fmt"
)

func main() {
	var (
		Kalkulator             string
		nilai1, nilai2, nilai3 int
	)

	fmt.Println(" Membuat Program Kalkulator")
	fmt.Println(" ")

	fmt.Println(" Pilihan Kalkulator \n Penjumlahan (+) \n Pengurangan (-) \n Perkalian (*), \n Pembagian (/) \n Modulus (%)")
	fmt.Print("Masukkan nilai pertama : ")
	fmt.Scan(&nilai1)
	fmt.Print("Masukkan nilai kedua : ")
	fmt.Scan(&nilai2)
	fmt.Print("Masukkan Kalkulator yang diinginkan : ")
	fmt.Scan(&Kalkulator)

	if Kalkulator == "Penjumlahan" || Kalkulator == "+" {
		nilai3 = nilai1 + nilai2
	} else if Kalkulator == "Pengurangan" || Kalkulator == "-" {
		nilai3 = nilai1 - nilai2
	} else if Kalkulator == "Perkalian" || Kalkulator == "*" {
		nilai3 = nilai1 * nilai2
	} else if Kalkulator == "Pembagian" || Kalkulator == "/" {
		nilai3 = nilai1 / nilai2
	} else if Kalkulator == "Modulus" || Kalkulator == "%" {
		nilai3 = nilai1 % nilai2
	} else {
		fmt.Println("Pilihan Kalkulator tidak tersedia")
	}
	fmt.Println("Hasil dari ", nilai1, Kalkulator, nilai2, " = ", nilai3)
}
