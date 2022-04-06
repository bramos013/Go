package main

import (
	"fmt"
	"time"
)

func notaParaConceito(nota float64) string {
	switch {
	case nota >= 9 && nota <= 10:
		return "A"
	case nota >= 8 && nota < 9:
		return "B"
	case nota >= 5 && nota < 8:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	case nota >= 0 && nota < 3:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	t := time.Now()
	switch { // switch sem argumento
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")

	}

	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(6.9))

}
