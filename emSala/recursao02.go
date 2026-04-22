package main

import "fmt"

func somar(arr []int) int{
	if len(arr) == 0{
		return 0
	}
	return arr[0] + somar(arr[1:]) //arr[1:] cria um slice de array que contém todos os elementos do array original, menos o de índice 0
}

func main() {
	const tamanho = 5
	array := []int{10, 20, 30, 40, 50}
	resultado := somar(array)
	fmt.Printf("A soma total é: %v\n", resultado)
}