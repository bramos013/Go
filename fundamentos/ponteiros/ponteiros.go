package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereço da variável i
	*p++   // desreferenciando (pegando o valor da variável i)
	i++

	//Go não tem aritmética de ponteiros
	//p++

	fmt.Println(p, *p, i)
}
