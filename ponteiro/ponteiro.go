package main

import "fmt"

func main() {
	i := 1

	var p *int = nil // Criando um ponteiro nulo

	p = &i // referênciando esse ponteiro para o endereço da variável (i)

	*p++ // Buscando o valor para onde o ponteiro está apontando e incrementando
	i++  // Incrementando diretamente o valor da varivel referenciada por (*p)

	// Go não tem aritimética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}
