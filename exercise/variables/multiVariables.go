package main

import "fmt"

func main() {
	a := "Green"
	fmt.Println("My favourite colour is ", a)

	b, c := 1984, 39
	fmt.Println("Born in ", b, "aged ", c)

	var (
		d = "R"
		e = "P"
		f = "G"
	)
	fmt.Println("The game is ", d, e, f)

	var g int
	g = 365 * g
	fmt.Println("The baby is ", g, " days old.")
}
