package main

import (
	"fmt"     // formatação de entrada e saída
	"math"    // para manipulação de funções aritiméticas
	"reflect" // Para averiguação de tipos
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200)) // TypeOf 3200 é um INT64
	var b byte = 255                                       // byte é o menor tipo numérico
	fmt.Println("O byte é", reflect.TypeOf(b))             // TypeOf b é : byte

	i1 := math.MaxInt64 // O maior número possível na estrutura Int64
	fmt.Println("O valor máximo é:", i1)
	fmt.Println("O tipo do valor é:", reflect.TypeOf(i1))

	var i2 rune = 'a' // Representa um mapeamento da tabela unicode
	fmt.Println("O bitcode correspondente a i2 é:", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Float64 e Float32
	var x = float32(49.99) // Uma forma de mudar o tipo do literal
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo) // Negação do tipo bool true = false

	//String
	s1 := "olá meu nome é Marcos"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// String de múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Marcos`
	fmt.Println(s2)
	fmt.Println(len(s2))

	// char???
	char := 'a' // char é um INT32 - Mapeado pela tabela unicode
	fmt.Println("O tipo do char é:", reflect.TypeOf(char))
	fmt.Println(char)
}
