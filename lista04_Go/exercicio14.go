package main

import "fmt"

func main() {
	const tamanho int = 10
	var vetorA [tamanho]int 
	var vetorB [tamanho]int
	var vetorR [tamanho*2]int
	fmt.Println("Digite os termos do vetor A:")
	for i := 0; i < tamanho; i ++{
		fmt.Println(i + 1,"° termo:")
		fmt.Scan(&vetorA[i])
	}
	fmt.Println("Digite os termos do vetor B:")
	for i := 0; i < tamanho; i ++{
		fmt.Println(i + 1,"° termo:")
		fmt.Scan(&vetorB[i])
	}
	for i := 0; i < tamanho; i ++{
		vetorR[i*2] = vetorA[i]
		vetorR[i*2 + 1] = vetorB[i]
	}
	fmt.Print("[ ")
	for _, valor := range vetorR{
		fmt.Printf("%d ", valor)
	}
	fmt.Println("]")
}