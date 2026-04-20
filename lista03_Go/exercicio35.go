package main

import "fmt"

func main() {
	var n int
	fmt.Println("--ESTE PROGRAMA IMPRIME O EQUIVALENTE BINÁRIO DE NÚMEROS NO SISTEMA DECIMAL DE NUMERAÇÃO--")
	fmt.Println("Digite um número inteiro pertencente ao sistema decimal de numeração:")
	fmt.Scan(&n)
    fmt.Printf("Decimal: %d, Binário: %b\n", n, n) 
}