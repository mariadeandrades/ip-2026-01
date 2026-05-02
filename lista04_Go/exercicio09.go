package main

import "fmt"

func main() {
	array := make([]float64, 10)
	fmt.Println("Digite a altura de dez atletas a aseguir:")
	soma := 0.0
	for i := 0; i < 10; i ++{
		fmt.Println("Digiite a altura do", i + 1 ,"° atleta:")
		fmt.Scan(&array[i])
		soma += array[i]
	}
	media := soma/10
	novo := []float64 {}
	for i := 0; i < 10; i ++{
		if array[i] > media{
			novo = append(novo, array[i])
		}
	}
	fmt.Printf("Alturas acima da média: %v", novo)
}
