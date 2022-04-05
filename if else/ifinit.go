package main

import (
	"fmt"
	"math/rand"
	"time"
)

func aleatoriamente() int {
	s := rand.NewSource(time.Now().UnixNano()) // NewSource serve para parametrizar a randomização
	r := rand.New(s)                           // New inicializa a randomização com o parâmetro escolhido
	return r.Intn(10)                          // utilizamos a randomização criada com o parâmetro limite
}

func main() {
	if i := aleatoriamente(); i > 5 { // É possível inicializar uma variável no bloco IF
		fmt.Println("Você ganhou com o número:", i)
	} else {
		fmt.Println("Você perdeu com o número:", i)
	}
}
