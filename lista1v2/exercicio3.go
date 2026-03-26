package main 

import "fmt"

var n1, n2, n3 int32

func main() {
	fmt.Println("Digite o primeiro número da composição, entre 1 e 9.")
	fmt.Scan(&n1)
	for n1 == 0 || n1 > 9{
		fmt.Println("Número inválido. Digite um número entre 1 e 9.")
	}
	fmt.Println("Digite o segundo número da composição, entre 1 e 9.")
	fmt.Scan(&n2)
	for n2 == 0 || n2 > 9{
		fmt.Println("Número inválido. Digite um número entre 1 e 9.")
	}
	fmt.Println("Digite o terceiro número da composição, entre 1 e 9.")
	fmt.Scan(&n3)
	for n2 == 0 || n2 > 9{
		fmt.Println("Número inválido. Digite um número entre 1 e 9.")
	}
	concatena := fmt.Sprintf("%d%d%d", n1, n2, n3)
	fmt.Println("A concatenação resulta em", concatena)
}

