package main 

import"fmt"

func main() {
	var n int
	fmt.Println("Digite um número inteiro de três algarismos.")
	fmt.Scan(&n)
	for n < 100 || n > 999{
		fmt.Println("Número inválido. Digite um número de três dígitos.")
		fmt.Scan(&n)
	}
	dezena := ((n/10)%10)
	fmt.Printf("A dezena do número %v é %v.", n, dezena)
}
 