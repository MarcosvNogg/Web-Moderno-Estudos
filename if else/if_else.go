package main

import "fmt"

func avaliarNota(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota:", nota)
	} else {
		fmt.Println("Reprovado com nota:", nota)
	}
}

func main() {
	avaliarNota(5)
	avaliarNota(9)
}
