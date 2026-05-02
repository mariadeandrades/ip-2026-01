package main

import "fmt"

func main() {
	const tamanho int = 10
	array := make([]int, tamanho)
	impares := []int {}
	pares := []int {}
	var soma int = 0
	contador := 0
	fmt.Println("Digite, a seguir, 10 números para compor o array:")
	for i := 0; i < tamanho; i ++{
		fmt.Println("Digite o", i + 1,"° número.")
		fmt.Scan(&array[i])
		if array[i] % 2 != 0{
			impares = append(impares, array[i])
			contador += 1
		}else{
			pares = append(pares, array[i])
		}
	}
	for j := 0; j < len(pares) ; j ++{
		soma += pares[j]
	}
	fmt.Printf("Os números pares são: %v\nA soma dos números pares é %v\nOs números ímpares são: %v\nHá %v números ímpares\n", pares, soma, impares, contador)
}