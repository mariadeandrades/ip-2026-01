package main

import "fmt"

func main() {
	var vetor [10]float64
	var codigo int
	fmt.Println("Digite 10 números reais para preencher o vetor:")
	for i := 0; i < 10; i++ {
		fmt.Printf("Posição [%d]: ", i)
		fmt.Scan(&vetor[i])
	}
	for {
		fmt.Print("\nDigite o código (0:Sair, 1:Direta, 2:Inversa): ")
		fmt.Scan(&codigo)
		if codigo == 0 {
			fmt.Println("Programa encerrado.")
			break
		}
		switch codigo {
		case 1:
			fmt.Println("Ordem Direta:")
			for i := 0; i < 10; i++ {
				fmt.Printf("%.2f ", vetor[i])
			}
			fmt.Println()
		case 2:
			fmt.Println("Ordem Inversa:")
			for i := 9; i >= 0; i-- {
				fmt.Printf("%.2f ", vetor[i])
			}
			fmt.Println()
		default:
			fmt.Println("Código inválido! Tente novamente.")
		}
	}
}