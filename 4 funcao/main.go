package main

func main() {
	println(calcular(12, 23))

	f := notas

	resultado := f(1, 2, 3, 4)
	println(resultado)

	println("calculadora")
	soma, subtracao := caculadora(10, 59)
	println(soma)
	println(subtracao)

	funcao_1, funacao_2 := banco()

	println(funcao_1(500))
	println(funacao_2(600))

	println("pagar e receber")
	pagamento, recebimento := pagarReceber(10, 30)
	println(pagamento)
	println(recebimento)
}

func banco() (func(valor int64) int64, func(valor int64) int64) {
	funcao_pagar := pagar
	funacao_receber := receber

	return funcao_pagar, funacao_receber
}

func pagarReceber(valor1, valor2 int64) (int64, int64) {
	return pagar(valor1), receber(valor2)
}

func pagar(valor int64) int64 {
	return valor
}

func receber(valor int64) int64 {
	return valor
}

func notas(n1, n2, n3, n4 int16) int16 {
	return n1 + n2 + n3 + n4
}

func calcular(n1, n2 int64) int64 {
	return n1 + n2
}

func caculadora(n1, n2 int64) (int64, int64) {
	return adicao(n1, n2), subtracao(n1, n2)
}

func adicao(n1, n2 int64) int64 {
	return n1 + n2
}

func subtracao(n1, n2 int64) int64 {
	return n1 - n2
}
