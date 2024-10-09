package main

import "fmt"

func main() {
	var C, R, F, K float64

	fmt.Println("Temperatur dalam satuan Celcius:")
	fmt.Scan(&C)

	R = (C * 4 / 5)
	F = (R * 9 / 4) + 32
	K = (F + 459.67) * 5 / 9

	fmt.Println("Temperatur Celcius:", C)
	fmt.Println("Temperatur Reamur:", R)
	fmt.Println("Temperatur Fahrenheit", F)
	fmt.Printf("Temperatur Kelvin %.0f", K)
}
