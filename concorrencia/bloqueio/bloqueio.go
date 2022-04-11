package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante (espera pelo valor) - envio de dados
	fmt.Println("Só depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operacao bloqueante (espera pelo valor) - recebimento de dados
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock - não foi possível ler pois o canal está vazio
	fmt.Println("Fim")
}
