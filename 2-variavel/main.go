package main

import (
	"fmt"
)

func main() {
	//tipos variavel
	var nome string = "andre"
	sobrenome := "andre"
	idade := 18

	fmt.Println(nome)
	fmt.Println(sobrenome)
	fmt.Println(idade)

	var (
		endereco = "AV"
		rua      = "154"
	)
	print(endereco, rua)
	banco, conta, digito := "2020", "1010", "5959"
	fmt.Println(banco, conta, digito)

	const URL = "www.google.com.br"

}
