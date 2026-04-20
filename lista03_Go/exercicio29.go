package main

import "fmt" 

func main() {
	var N int
	fmt.Println("--ESTE PROGRAMA CALCULA A SOMA DOS N PRIMEIROS TERMOS, DE 1 A N--")
	fmt.Println("Digite o N de sua escolha:")
	fmt.Scan(&N)
	for N < 0{
		fmt.Println("Digite somente valores positivos para N:")
		fmt.Scan(&N)
	}
	soma := 0
	for i := 1; i <= N; i ++{
		soma += i
	}
	fmt.Printf("O valor da soma é %v.", soma)
}