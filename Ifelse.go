package main

import (
	"fmt"
)

func main() {
	var (
		celcius, Suhu float64
		Konversi      string
	)
	fmt.Println(" Pilih Konversi yang dari celcius ke  : \n1. Fahrenheit \n2. Reamur \n3. Kelvin")
	fmt.Print("Masukkan konversi : ")
	fmt.Scan(&Konversi)
	fmt.Print("Masukkan suhu dalam derajat celcius : ")
	fmt.Scan(&celcius)
	if Konversi == "1" || Konversi == "Fahrenheit" {
		Suhu = ((9 / 5 * celcius) + 32)
	} else if Konversi == "2" || Konversi == "Reamur" {
		Suhu = (4 / 5) * celcius
	} else if Konversi == "3" || Konversi == "Kelvin" {
		Suhu = celcius + 273.15
	} else {
		fmt.Println("Konversi Tidak Tersedia")
	}
	fmt.Println("Hasil Konversi Suhu Celcius ke", Konversi, "adalah ", Suhu)

}
