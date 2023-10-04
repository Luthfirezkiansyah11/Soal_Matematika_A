package main

import "fmt"

func main() {

	var x, y float64

	fmt.Print("masukan:")
	fmt.Scanln(&x, &y)
	var hasil = (1/(3*(x*x)+10) + (10 * y) + 7)
	fmt.Println(hasil)

}
