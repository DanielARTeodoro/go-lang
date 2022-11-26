package main

import "fmt"

func endentificacaoDoAluno() {

	fmt.Println("Coloque o Primeirro Nome do Aluno: ")

	var first string

	fmt.Scanln(&first)
	fmt.Println("Coloque o Utimo Nome do Aluno: ")
	var second string
	fmt.Scanln(&second)

	fmt.Print("O Nome completo do aluno a ser Comfirmado é: ")
	fmt.Print(first + " " + second)

}

func main() {

	endentificacaoDoAluno()

}
