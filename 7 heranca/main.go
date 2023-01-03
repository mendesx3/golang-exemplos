package main

import "fmt"

type pessoa struct {
	nome  string
	idade int16
	sexo  string
}

type estudante struct {
	pessoa
	ra    string
	curso string
}

func main() {

	e1 := estudante{pessoa: pessoa{nome: "andre", idade: 19, sexo: "M"}, ra: "1441", curso: "TI"}
	fmt.Println(e1)

	p1 := pessoa{"andre", 38, "m"}

	e2 := estudante{pessoa: p1, ra: "444", curso: "TI"}
	fmt.Println(e2)
	fmt.Println("aluno", e2.nome)
	fmt.Println("curso", e2.curso)
	fmt.Println("dados completo", e2.pessoa)

}
