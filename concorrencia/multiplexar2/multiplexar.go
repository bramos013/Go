package main

import (
	"fmt"
	"time"
)

func falar(pessoa string) <-chan string {
	c := make(chan string) // canal sem buffer
	go func() {            // função anônima
		for i := 0; i < 3; i++ { // repetição
			c <- fmt.Sprintf("%s falando: %d", pessoa, i) // envia dados para o canal
			time.Sleep(time.Second)                       // pausa
		}
	}() // função anônima
	return c
}
func juntar(entrada1, entrada2 <-chan string) <-chan string { // recebe dois canais de strings
	c := make(chan string) // canal sem buffer
	go func() {            // função anônima
		for { // repetição
			select {
			case s := <-entrada1: // recebe dados do canal 1
				c <- s // envia dados para o canal
			case s := <-entrada2:
				c <- s
			}
		}
	}() // função anônima
	return c
}

func main() {
	c := juntar(falar("João"), falar("Maria")) // junta dois canais de strings
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
