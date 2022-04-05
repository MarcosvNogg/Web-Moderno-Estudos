package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 7 && nota < 8 {
		return "C"
	} else if nota >= 4 && nota < 7 {
		return "D"
	} else if nota >= 2 && nota < 4 {
		return "E"
	} else {
		return "F"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(8))
	fmt.Println(notaParaConceito(7.1))
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(2.9))
	fmt.Println(notaParaConceito(1))
}
