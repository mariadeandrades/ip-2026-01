package main 

import "fmt"

func main() {
	var soma float64

	fmt.Println("Este programa realiza a seguinte opreação matemática:")
	fmt.Println("H = 1/1 + 3/2 + 5/3 + 7/4 + ... + 99/50")

	//numerador = n | denominador = d
	n := 0.0
	for d := 1.0; d <= 50; d ++{
		termo := (n/d)
		soma += termo
		fmt.Printf("Termo: %.0f/%.0f = %.2f\n", n, d, termo)
		n += 2
    }
	fmt.Printf("H = %.2f", soma)
}