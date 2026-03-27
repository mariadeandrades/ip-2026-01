package main

import "fmt"

var a, b, c, D float32

func main() {
	fmt.Println("Qual o valor dos coeficientes a, b e c?")
	fmt.Scan(&a, &b, &c)
	D := (b*b - 4*a*c)
	fmt.Println("O valor de delta é", D)
}