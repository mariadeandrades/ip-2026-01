package main 

import "fmt"

func main() {
	var n int

	fmt.Println("--ESTE PROGRAMA CALCULA O FATORIAL DE NÚMEROS!--")
	fmt.Println("INSTRUÇÕES GERAIS: Não existe fatorial de números negativos (menores que zero);\npor isso, digite apenas números inteiros e positivos (maiores que zero)")
	fmt.Println("Digite um número para calcular o fatorial:")
	fmt.Scan(&n)

	for n < 0{
		fmt.Println("Digite apenas números inteiros positivos (maiores que zero):")
		fmt.Scan(&n)
	}

	resultado := uint(1)
	for i := 1; i <= n; i ++{  //em loops de multiplicação, começar o i em 1 para não gerar sempre respostas = 0
			resultado *= uint(i)
		}
	fmt.Printf("%v! = %v", n, resultado)
}