package main

/// RAILSON 9 - Receber três números e faça a multiplicação entre eles

import "fmt"

var (
	x  int
	n  int
	n0 int
)

func multiplicacao() {

	fmt.Println("Digite três valores: ")
	fmt.Println("Digite o valor 1: ")
	fmt.Scan(&n)
	fmt.Println("Digite o valor 2: ")
	fmt.Scan(&n0)
	fmt.Println("Digite o valor 3: ")
	fmt.Scan(&x)

	fmt.Println("A Mutiplicação é: ", n*n0*x)

}

func main() {

	multiplicacao()

}
