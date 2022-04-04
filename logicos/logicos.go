package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	Tv50 := trab1 && trab2
	Tv32 := trab1 != trab2
	Sorvete := trab1 || trab2
	return Tv50, Tv32, Sorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
