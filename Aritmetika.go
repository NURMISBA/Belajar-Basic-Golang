package main

import (
	"fmt"
)

func main() {
	var (
		sisi   int
		jari   int
		alas   int
		tinggi int
	)
	const phi float32 = 3.14

	fmt.Println("Menghitung Luas Segitiga ")
	alas = 5
	tinggi = 6
	var luas_segitiga float32 = 0.5 * float32(alas) * float32(tinggi)
	fmt.Println("Luas segitiga dari alas = ", alas, "dan tinggi = ", tinggi, "yang dimana rumusnya adalah 1/2 * alas * tinggi = ", luas_segitiga, "cm")
	fmt.Println(" ")

	fmt.Println("Menghitung Luas Lingkaran ")
	jari = 8
	luas_lingkaran := phi * float32(jari) * float32(jari)
	fmt.Println("Luas lingkaran dari jari-jari = ", jari, " yang dimana rumusnya adalah phi * jari-jari * jari-jari = ", luas_lingkaran, "cm")
	fmt.Println(" ")

	fmt.Println("Menghitung Luas persegi ")
	sisi = 7
	luas_persegi := sisi * sisi
	fmt.Println("Luas persegi dari sisi persegi =", sisi, " yang dimana rumusnya adalah sisi * sisi = ", luas_persegi, "cm")
	fmt.Println(" ")
}
