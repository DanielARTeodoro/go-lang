package main

/// GEOVANY 14 - crie uma função que calcula o IMC

import "fmt"

var (
	imc    float64
	peso   float64
	altura float64
)

func imccalc() {
	fmt.Println("Digite o seu IMC: ")
	fmt.Println("Digite o seu Peso: ")
	fmt.Scan(&peso)
	fmt.Println("Digite a sua Altura: ")
	fmt.Scan(&altura)
	imc = (peso) / (altura * altura)
	fmt.Printf("O seu IMC é: %0.01f", imc)
}

func main() {

	imccalc()

}
