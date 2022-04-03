package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	//cuidado...
	fmt.Println("teste: " + fmt.Sprint(123)) // Podemos usar Sprint para converter Inteiro em String

	// int para String...
	fmt.Println("Teste:" + strconv.Itoa(123)) // Teste: "123" - Itoa converte inteiro para String

	// string para int
	num, _ := strconv.Atoi("123") // Atoi converte uma string em um número inteiro
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("true") // true e 1 são valores verdadeiros, o resto é falso
	if b {
		fmt.Println("Verdadeiro")
	}

}
