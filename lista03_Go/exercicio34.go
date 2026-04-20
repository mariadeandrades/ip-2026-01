package main

import "fmt"

//a lógica aqui é usar o algoritmo Euclidiano
func calcularMDC(a, b int) int{
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func calcularMMC(a, b int) int{
	if a == 0 || b == 0 {
		return 0
	}
	return (a * b) / calcularMDC(a, b)
}

func main() {
	var n1, n2 int

	fmt.Println("-- CALCULADORA DE MMC --")
	fmt.Print("Digite o primeiro número (N1): ")
	fmt.Scan(&n1)
	fmt.Print("Digite o segundo número (N2): ")
	fmt.Scan(&n2)

	if n1 <= 0 || n2 <= 0{
		fmt.Println("Por favor, digite apenas números inteiros positivos.")
		return
	}
	resultado := calcularMMC(n1, n2)
	fmt.Printf("O MMC de %d e %d é: %d\n", n1, n2, resultado)
}