package main

/// ANTONIA 6 - Uma função que cria uma calculadora simples, digitando o n1, n2 e a operação desejada

import (
	"fmt"
)

var (
	x  int
	n  int
	n0 int
)

func calculadoraSimples() {

	fmt.Println("Esconlha uma Operação Simples: ")
	fmt.Println("1 -> 1 +1 , 2 -> 1 - 1, 3 -> 1 x 1, 4 -> 1 / 1: ")
	fmt.Scan(&x)

	switch x {
	case 1:
		soma()
	case 2:
		subtracao()
	case 3:
		multiplicacao()
	case 4:
		divisao()
	}
}

func soma() {

	fmt.Println("Digite dois valores: ")
	fmt.Println("Digite o valor 1: ")
	fmt.Scan(&n)
	fmt.Println("Digite o valor 2: ")
	fmt.Scan(&n0)

	fmt.Println("A Soma é: ", n+n0)

}

func subtracao() {

	fmt.Println("Digite dois valores: ")
	fmt.Println("Digite o valor 1: ")
	fmt.Scan(&n)
	fmt.Println("Digite o valor 2: ")
	fmt.Scan(&n0)

	fmt.Println("A Subtração é: ", n-n0)

}

func multiplicacao() {

	fmt.Println("Digite dois valores: ")
	fmt.Println("Digite o valor 1: ")
	fmt.Scan(&n)
	fmt.Println("Digite o valor 2: ")
	fmt.Scan(&n0)

	fmt.Println("A Mutiplicação é: ", n*n0)

}

func divisao() {

	fmt.Println("Digite dois valores: ")
	fmt.Println("Digite o valor 1: ")
	fmt.Scan(&n)
	fmt.Println("Digite o valor 2: ")
	fmt.Scan(&n0)

	fmt.Println("A Divisão é: ", n/n0)

}

func main() {
	calculadoraSimples()
}
