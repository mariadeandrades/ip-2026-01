package main

import "fmt"

func main() {
	const tamanho = 50
	var idades [tamanho]int
	contagem := make(map[int]int)
	fmt.Printf("Digite as %d idades:\n", tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Printf("%dª idade: ", i+1)
		fmt.Scan(&idades[i])
		contagem[idades[i]]++
	}
	moda := idades[0]
	maxRepeticoes := 0
	for idade, qtd := range contagem {
		if qtd > maxRepeticoes {
			maxRepeticoes = qtd
			moda = idade
		}
	}
	fmt.Printf("\n--- Resultado ---\n")
	if maxRepeticoes == 1 {
		fmt.Println("Não existe moda (amodal): todas as idades aparecem apenas uma vez.")
	} else {
		fmt.Printf("A moda é: %d anos (aparece %d vezes).\n", moda, maxRepeticoes)
	}
}