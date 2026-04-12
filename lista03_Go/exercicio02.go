package main

import "fmt"

func main() {
	var soma int
	var quantidade int
	var media float64

	for i := 50; i <= 70; i += 2 {
		soma += i
		quantidade++  //quantidade = quantidade + 1
	}

	media = float64(soma) / float64(quantidade)

	fmt.Printf("A soma dos números pares entre 50 e 70 é: %d\n", soma)
	fmt.Printf("A média é: %.2f\n", media)
}