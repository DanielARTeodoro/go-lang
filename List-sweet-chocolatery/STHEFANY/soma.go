package main

/// STHEFANY 4 - crie uma função que some duas notas
import "fmt"

func soma() {

	var (
		n  int
		n0 int
	)

	fmt.Println("Digite dois valores: ")
	fmt.Println("Digite o valor 1: ")
	fmt.Scan(&n)
	fmt.Println("Digite o valor 2: ")
	fmt.Scan(&n0)

	fmt.Println("A soma é: ", n+n0)

}
func main() {
	soma()
}
