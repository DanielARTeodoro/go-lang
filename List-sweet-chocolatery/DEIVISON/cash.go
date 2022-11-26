package main

/// DEIVISON 13 - Criar uma função que converta a moeda Real , para moeda dólar

import "fmt"

var (
	real  float64
	dolar float64
)

func converterMoeda() {

	fmt.Println("Digite o valor do Real: ")
	fmt.Scan(&real)
	fmt.Println("Digite o valor da cotação do Dolar para o Real: ")
	fmt.Scan(&dolar)

	fmt.Printf("O valor em Dolar é: %0.01f", real/dolar)

}

func main() {

	converterMoeda()

}
