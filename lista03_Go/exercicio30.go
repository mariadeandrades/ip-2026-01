package main

import "fmt"

func main() {
	var R float64
	const pi float64 = 3.14
	fmt.Println("--ESTE PROGRAMA CALCULA O VOLUME DE UAM ESFERA QUALQUER--")
	fmt.Println("Para calcular o volume de uma esfera, digite o valor do seu raio:")
	fmt.Println("Obs: digite somente valores múltiplos de 0.5 para o raio da esfera.")
	fmt.Scan(&R)
	volume := ((4*pi*R*R*R)/3)
	fmt.Printf("O volume da esfera de raio %v é %v.", R, volume)
}