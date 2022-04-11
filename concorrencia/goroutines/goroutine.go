//Goroutines fundamento de concorrencia em Go
package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second) //simula um atraso para ver a concorrencia
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}
func main() {
	//("Maria", "Por que você não fala comigo?", 3)
	//fale("João", "Só posso falar depois de você!", 1)

	//criando goroutines utilizando a palavra-chave go
	// Nesse formato a goroutine é iniciada imediatamente
	// e a função fale() é executada em uma nova goroutine
	// Uma goroutine é uma thread em Go
	// Ela só irá funcionar se a thread principal estiver executando
	//go fale("Maria", "Ei...", 500)
	//go fale("João", "Opa...", 500)
	// Se nao colocar um tempo de espera, a goroutine principal irá morrer
	// e a função fale() não será executada
	//time.Sleep(time.Second * 7)

	//Outro exemplo
	go fale("Maria", "Entendi!!!", 10)
	fale("João", "Parabéns!!!", 5)
}
