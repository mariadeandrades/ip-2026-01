package main

import (
	"fmt"
	"math"
)

func main() {
	const tamanho = 100
	var b [tamanho]float64
	var soma float64

	fmt.Printf("Digite os %d números:\n", tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Printf("Valor %d: ", i)
		fmt.Scan(&b[i])
	}
	for i := 0; i <= 49; i++ {
		diferenca := b[i] - b[99-i]
		soma += math.Pow(diferenca, 3)
	}
	fmt.Printf("\nO valor da soma S é: %.2f\n", soma)
}