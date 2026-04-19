package main

import "fmt"

func main() {
	soma := 0.0
	num := 100.0
	den := 1.0
	for i := 0; i <= 20; i++{
		if i > 0{
			den *= float64(i)
		}
		termo := num/den
		num -= 1
		soma += termo
	}
	fmt.Printf("O resultado da soma é %.2f.", soma)
}