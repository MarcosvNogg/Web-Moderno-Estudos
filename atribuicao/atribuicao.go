package main

import "fmt"

func main() {
	var b byte = 5 // Atribuição comum
	fmt.Println(b)

	i := 4 // Atribuição + Inicialização
	i += 1 // Atribuição Aditiva
	i -= 1 // Atribuição subtrativa
	i *= 1 // Atribuição Multiplicativa
	i /= 2 // Atribuição Divisiva
	i %= 2 // Atribuição de Módulo
	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	x, y = y, x
	fmt.Println(x, y) // Inversão de valores
}
