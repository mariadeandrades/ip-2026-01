package main

import "fmt"

func main() {
	const tamanho int = 30
	var vetorA [tamanho]int
	var vetorB [tamanho]int
	fmt.Println("A seguir, digite trinta valores para preencher o vetor A:")
	for i := 0; i < tamanho; i ++{
		fmt.Println(i + 1,"° valor:")
		fmt.Scan(&vetorA[i])
	}
	for i := 0; i < tamanho; i ++{
		if i % 2 == 0{
			vetorB[i] = 2*vetorA[i]
		}else{
			vetorB[i] = 3*vetorA[i]
		}
	}
	fmt.Print("[")
	for _, valor := range vetorB{
		fmt.Printf(" %d ", valor)
	}
	fmt.Println("]")
}