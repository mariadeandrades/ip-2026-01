package main

import "fmt"

func main() {
	var N1, N2 int
	fmt.Println("--ESTE PROGRAMA CALCULA O RESTO E O QUOCIENTE DA DIVISÃO DE UM NÚMERO N1 POR UM NÚMERO N2\nUSANDO SOMENTE ADIÇÃO--")
	fmt.Println("Digite os números N1 e N2, sendo que o número N2 será o divisor de N1 (N1 / N2 = resultado):")
	fmt.Scan(&N1, &N2)
	soma := 0
	quociente := 0
	for soma + N2 < N1{
		soma += N2
		quociente++
	}
	resto := N1 - soma
	fmt.Printf("Ao dividir %v por %v, o quociente é %v e o resto é %v.", N1, N2, quociente, resto)
}