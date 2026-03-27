package main

import "fmt"

var F, P, C, M float32

func main() {
	fmt.Println("Digite o valor em Fahrenheit a ser convertido.")
	fmt.Scan(&F)
	C = (5*(F-32)/9)
	fmt.Println("O valor da temperatura em Celsius é", C)
	fmt.Println("Digite o valor em polegadas a ser convertido.")
	fmt.Scan(&P)
	M = (P*25.4)
	fmt.Println("O valor do volume de chuva em mm é", M)
}