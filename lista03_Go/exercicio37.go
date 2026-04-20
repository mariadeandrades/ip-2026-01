package main

import (
	"fmt"
	"math"
)

func main() {
	var octal int
	fmt.Print("Digite um número inteiro na base 8: ")
	fmt.Scan(&octal)
	decimal := 0
	potencia := 0
	numeroOriginal := octal
	for octal > 0 {
		digito := octal % 10
		decimal += digito * int(math.Pow(8, float64(potencia)))
		octal = octal / 10
		potencia++
	}
	fmt.Printf("O número octal %d em base 10 é: %d\n", numeroOriginal, decimal)
}