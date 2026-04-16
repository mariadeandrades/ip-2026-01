package main

import "fmt"

func fib() func() int {
	f1 := 0 //o primeiro e o segundo termos são 0 1 e precisam ser declarados antes
	f2 := 1
	return func() int {
			f1, f2 = f2, (f1 + f2)
			return f1
			}
}

func main() {
	var t int
	fmt.Println("--Sequência de Fibonacci--\nEste programa imprime a sequência de Fibonacci do tamanho que\nvocê quiser! Digite o número de termos (maior que 3) que deseja na sua sequência:")
	fmt.Scan(&t)
	fi := fib() 
	for i := 0; i < t; i ++{
		fmt.Printf("%d ", fi())
	}
}
/* 
0 1 1 2 3 5 8 13 21 35
a=0
b=1
a b a+b a+b+b a+b+a+b+b a+a+a+b+b+b+b+b
0 1 0+1 0+1+1 0+1+1+1	0+1+1+1+1+1
using recursion*/