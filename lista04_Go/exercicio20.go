package main

import "fmt"

func main() {
	const jogadas = 20
	var resultados [jogadas]int
	frequencia := make(map[int]int)
	fmt.Println("Digite os números sorteados (de 1 a 6) nas 20 jogadas:")
	for i := 0; i < jogadas; i++ {
		for {
			var num int
			fmt.Printf("%dª jogada: ", i+1)
			fmt.Scan(&num)
			if num >= 1 && num <= 6 {
				resultados[i] = num
				frequencia[num]++ 
				break
			} else {
				fmt.Println("Valor inválido! Por favor, digite um número entre 1 e 6.")
			}
		}
	}
	fmt.Println("\n--- Números Sorteados ---")
	fmt.Print("[ ")
	for _, v := range resultados {
		fmt.Printf("%d ", v)
	}
	fmt.Println("]")
	fmt.Println("\n--- Frequência das Faces ---")
	for face := 1; face <= 6; face++ {
		qtd := frequencia[face]
		fmt.Printf("Face %d: apareceu %d vez(es)\n", face, qtd)
	}
}