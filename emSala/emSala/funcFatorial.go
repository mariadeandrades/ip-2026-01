// aula do dia 01/04/2026
package main

import "fmt"

func fatorial(n uint) uint {
	if n == 0 {
		return 1
	}
	result := uint(1)
	for i := uint(1); i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	var n int
	fmt.Println("Digite um número inteiro não-negativo para calcular o fatorial: ")
	fmt.Scan(&n)
	for n < 0{
		fmt.Println("Entrada inválida, não é possível calcular o fatorial de um número negativo.\nPor favor, digite um número inteiro.")
		fmt.Scan(&n)
	}

	resultado := fatorial(uint(n))
	fmt.Printf("%d! = %d\n", n, resultado)
}