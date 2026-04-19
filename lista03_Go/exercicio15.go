package main

import "fmt"

func main() {
	var numeros[] int
	var termo, n int
	
	numeros = []int {}
	fmt.Println("--SEQUÊNCIA (a^2)*((a+1)^2)*...*((a+n-1)^2)")
	fmt.Println("Digite um número natural, que será o limite da sequência:")
	fmt.Scan(&n)
	for i := 1; i <= n; i ++{
		termo = i*i
		numeros = append(numeros, termo)
	}
	fmt.Printf("%v", numeros)
}