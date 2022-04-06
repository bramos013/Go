package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // não tem while no go
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando...")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// laço infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
