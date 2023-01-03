package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    int64
	endereco endereco
	pessoais dadosPessoais
}

type endereco struct {
	rua    string
	numero int64
}

type dadosPessoais struct {
	Rg  string `json:"rg"`
	Cpf string `json:"cpf"`
}

func newDadosPessoais(rg string, cpf string) *dadosPessoais {
	return &dadosPessoais{Rg: rg, Cpf: cpf}
}

type CRYPTO struct {
	uuid string
	name string
}

func main() {
	fmt.Println("aula structs")

	var c = CRYPTO{
		name: "andre",
		uuid: "1112",
	}

	fmt.Println(c)

	var u usuario
	u.nome = "andre"
	u.idade = 15
	fmt.Println(u)

	end := endereco{rua: "AV", numero: 1043}

	novoDados := newDadosPessoais("001", "2221")

	u2 := usuario{"andre", 44, end, *novoDados}
	fmt.Println(u2)

	u3 := usuario{idade: 44}
	fmt.Println(u3)
}
