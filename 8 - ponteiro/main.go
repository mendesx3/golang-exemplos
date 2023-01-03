package main

import "fmt"

func main() {
	fmt.Println("start ponteiro")

	var numero int64 = 10
	var numero2 = numero

	fmt.Println(numero)
	fmt.Println(numero2)

	numero++
	fmt.Println(numero)
	fmt.Println(numero2)

	var variavel3 = 100
	var ponteiro *int = &variavel3
	fmt.Println(variavel3)
	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)

}
