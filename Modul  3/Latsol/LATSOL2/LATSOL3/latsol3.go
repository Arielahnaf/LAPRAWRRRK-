package main

import "fmt"

func main() {
	var tahun int

	fmt.Println("tahun:")
	fmt.Scan(&tahun)
	fmt.Println()

	kabisat := (tahun%400 == 0) || (tahun%4 == 0) && (tahun%100 != 0)

	fmt.Println(tahun, "adalah tahun kabisat.")
	fmt.Println(kabisat)
}
