package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float32, float64:
		return "real"
	case func():
		return "funcao"
	default:
		return "outro"
	}
}

func main() {
	fmt.Println(tipo(2))
	fmt.Println(tipo("epa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(3.14))
	fmt.Println(tipo(true))
}
