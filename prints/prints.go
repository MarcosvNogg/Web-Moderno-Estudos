package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha")

	fmt.Println(" ... Nova")
	fmt.Println("Linha")

	x := 3.141516

	//fmt.Println("O valor de x é ")

	xs := fmt.Sprint(x)
	fmt.Println("A string formatada contem " + xs)
	fmt.Println("A sting formatada contem", x)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
}
