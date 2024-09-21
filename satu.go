package main

import (
	"fmt"
)

func main() {
	var celcius float64
	var pilihan int

	fmt.Print("Masukkan suhu dalam Celcius: ")
	fmt.Scan(&celcius)

	fmt.Println("Pilih konversi suhu:")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Kelvin")
	fmt.Println("3. Reamur")
	fmt.Print("Masukkan pilihan (1/2/3): ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		// Konversi ke Fahrenheit
		fahrenheit := (celcius * 9 / 5) + 32
		fmt.Printf("%.2f Celcius = %.2f Fahrenheit\n", celcius, fahrenheit)
	case 2:
		// Konversi ke Kelvin
		kelvin := celcius + 273.15
		fmt.Printf("%.2f Celcius = %.2f Kelvin\n", celcius, kelvin)
	case 3:
		// Konversi ke Reamur
		reamur := celcius * 4 / 5
		fmt.Printf("%.2f Celcius = %.2f Reamur\n", celcius, reamur)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}
