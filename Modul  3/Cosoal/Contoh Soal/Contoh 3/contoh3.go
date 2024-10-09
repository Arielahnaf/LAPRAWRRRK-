package main

import "fmt"

func main() {
	var rupiah, dolar float64
	fmt.Print("Masukan Nominal Rupiah: ")
	fmt.Scan(&rupiah)
	dolar = (rupiah / 15000)
	fmt.Print("Jadi ", rupiah, " rupiah = ", dolar, "dolar")
}
