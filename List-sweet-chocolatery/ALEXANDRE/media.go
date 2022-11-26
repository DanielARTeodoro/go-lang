package main

/// ALEXANDRE 5 - Criar uma função que faça média entre notas
import "fmt"

func media() {

	var (
		n     int
		n0    int
		n1    int
		n01   int
		media int
	)

	fmt.Println("Digite a Media das Notas: ")
	fmt.Println("Digite 4 valores: ")
	fmt.Println("Digite o valor 1: ")
	fmt.Scan(&n)
	fmt.Println("Digite o valor 2: ")
	fmt.Scan(&n0)
	fmt.Println("Digite o valor 3: ")
	fmt.Scan(&n1)
	fmt.Println("Digite o valor 4: ")
	fmt.Scan(&n01)
	media = n + n0 + n1 + n01
	fmt.Println("A media é: ", media/4)

}
func main() {
	media()
}
