package main

import "fmt"

func main() {

	//5 operadores relacionais
	fmt.Println(1 > 0)
	fmt.Println(1 > 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 == 1)
	fmt.Println(1 != 1)

	//5 operadores logicos
	if "andre" == "joao" {
		println("é igual")
	}

	if "andre" == "andre" && 1 == 1 {
		println("é igual")
	}

	numero := 10
	println(numero)
	numero++
	println(numero)
	numero += 20
	println(numero)

}
