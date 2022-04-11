package main

import "fmt"

func main() {
	// Canal é um tipo especial de variável que permite a comunicação entre goroutines
	ch := make(chan int, 1)

	ch <- 1 //enviando dados para o canal(escrita)
	<-ch    //recebendo dados do canal(leitura)

	ch <- 2           //enviando dados para o canal(escrita)
	fmt.Println(<-ch) //recebendo dados do canal(leitura)

}
