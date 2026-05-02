/*Faça um programa que leia um vetor de inteiros, de 10 posições. A seguir, encontre o menor elemento (X) do
vetor. Imprima uma mensagem mostrando: “O menor elemento do vetor é”, X, “e sua posição dentro do vetor
é:”,P. Assuma que os elementos informados no vetor são todos diferentes entre si.*/

package main

import "fmt"

func main() {
	array := make([]int, 10)
	fmt.Println("Digite 10 números para preencher o array:")
	for i := 0; i < 10; i ++{
		fmt.Println("Digite o", i + 1 ,"° valor:")
		fmt.Scan(&array[i])
	}
	min := array[0]
	position := 0
	for j := 0; j < 10; j ++{
		if array[j] < min{
			min = array[j]
			position = j
		}
	}
	fmt.Printf("O menor elemento do vetor é %d e sua posição dentro do vetor é: %d\n", min, position)
}
