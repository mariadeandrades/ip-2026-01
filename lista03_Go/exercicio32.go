package main

import "fmt"

func main() {
	var N1, N2 int
	fmt.Println("--ESTE PROGRAMA CALCULA O PRODUTO DE UM NÚMERO N1 POR UM NÚMERO N2 USANDO ADIÇÃO--")
	fmt.Println("Digite os números N1 e N2, sendo que N1 será o multiplicador de N2 (N1 * N2 = resultado):")
	fmt.Scan(&N1, &N2)
	resultado := 0
	for i := 1; i <= N1; i ++{
		resultado += N2
	}
	fmt.Printf("O resultado da soma entre %v e %v, usando adição, é %v.", N1, N2, resultado)
}