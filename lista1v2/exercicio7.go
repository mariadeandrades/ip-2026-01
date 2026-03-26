/*
algoritmo_conversoes_gerais

declare F, P float

escreva ("Digite o valor em Fahrenheit a ser convertido.")
receba F
funcao C(F)
    retorne ((5F - 160)/9)
escreva ("O valor da temperatura em Celsius é", C)
escreva ("Digite o valor em polegadas a ser convertido.")
receba P
funcao M(P)
    retorne (P*25.4)
escreval ("O valor do volume de chuva em mm é", M)

fim_algoritmo_conversoes_gerais
*/ 

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