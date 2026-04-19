package main

import "fmt"

func main() {
	var N int

	fmt.Println("Digite um número natural N, que será o tamanho da sequência:")
	fmt.Scan(&N)
	soma := 0.0
	n := 1000.0
	d := 0.0
	for i := 1; i <= N; i ++{
		soma = n/d
		if i%2 == 0{  //essa parte faz a alternância de sinais
			soma -= termo
		}else{
			soma += termo
		}
		n -= 3
		d += 1
	}
	fmt.Printf("A soma é %v", soma)
}