package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	// Realiza a operação de envio de dados para o canal com buffer de tamanho 3
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executando rotina...") //Executando rotina...
	ch <- 5
	ch <- 6
}

func main() {
	ch := make(chan int, 3) //Canal com buffer de tamanho 3
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
