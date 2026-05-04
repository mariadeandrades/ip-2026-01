package main

import "fmt"

func main() {
	const tamanho = 10
	var vetor [tamanho]int
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Posição [%d]: ", i)
		fmt.Scan(&vetor[i])
	}
	fmt.Println("\n--- Números Primos Encontrados ---")
	for i := 0; i < tamanho; i++ {
		numero := vetor[i]
		if ehPrimo(numero) {
			fmt.Printf("Número %d encontrado na posição %d\n", numero, i)
		}
	}
}
func ehPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}