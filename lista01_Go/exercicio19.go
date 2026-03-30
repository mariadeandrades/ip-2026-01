package main 

import "fmt"

func main() {
	var n, i, soma int
	fmt.Println("Digite um número inteiro e positivo maior que 1.")
	fmt.Scan(&n)
	for n < 1{
		fmt.Println("Digite um número válido.")
		fmt.Scan(&n)
	}
	for i = 1; i <= n; i ++{
		soma = (soma + 1/i)
	}
	fmt.Printf("A soma é %d", soma)
}
// compilou mas não funciona, não sei como corrigir