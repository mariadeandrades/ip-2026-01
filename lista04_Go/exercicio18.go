package main

import "fmt"

func main() {
	const tamanho = 10
	var vetor [tamanho]int
	var preenchidos int 
	fmt.Println("Digite 10 números (eles serão ordenados automaticamente):")
	for preenchidos < tamanho {
		var novoNumero int
		fmt.Printf("%dº número: ", preenchidos+1)
		fmt.Scan(&novoNumero)
		posicao := 0
		for posicao < preenchidos && vetor[posicao] < novoNumero {
			posicao++
		}
		for j := preenchidos; j > posicao; j-- {
			vetor[j] = vetor[j-1]
		}
		vetor[posicao] = novoNumero
		preenchidos++
	}
	fmt.Println("\nVetor lido e armazenado de forma ordenada:")
	fmt.Printf("%v\n", vetor)
}