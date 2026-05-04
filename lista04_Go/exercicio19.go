package main

import "fmt"

func main() {
	var vetor1 [10]int
	var vetor2 [5]int

	// 1. Leitura do primeiro vetor (10 números)
	fmt.Println("Digite 10 números para o primeiro vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Vetor1[%d]: ", i)
		fmt.Scan(&vetor1[i])
	}
	fmt.Println("\nDigite 5 números para o segundo vetor:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Vetor2[%d]: ", i)
		fmt.Scan(&vetor2[i])
	}
	fmt.Println("\n--- Análise de Divisores ---")
	for i := 0; i < 10; i++ {
		num1 := vetor1[i]
		fmt.Printf("Número %d (posição %d):\n", num1, i)
		encontrouDivisor := false
		for j := 0; j < 5; j++ {
			num2 := vetor2[j]
			if num2 != 0 && num1%num2 == 0 {
				fmt.Printf("  -> Divisor: %d (no Vetor2, posição %d)\n", num2, j)
				encontrouDivisor = true
			}
		}
		if !encontrouDivisor {
			fmt.Println("  -> Nenhum divisor encontrado no segundo vetor.")
		}
		fmt.Println("---------------------------")
	}
}