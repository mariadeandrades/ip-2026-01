package main

import "fmt"

var h, a, v float64

func main() {
	fmt.Println("Digite a altura e a aresta da base da pirâmide, em metros.")
	fmt.Scan(&h, &a)
	v = ((a*a*3*3*h)/2)
	fmt.Println("O volume da pirâmide é %d metros cúbicos.", v)
}