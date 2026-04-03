package main 

import "fmt"

func teste(n int) int{
	return (8 / (2 - n))
}

func main() {
	var n, resultado int
	fmt.Println("Digite um número qualquer.")
	fmt.Scan(&n)
	resultado = teste(n)
	fmt.Printf("O resultado da opreção é %v.", resultado)
}