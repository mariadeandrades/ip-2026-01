package main

import "fmt"

func main() {
	var numeros []int
	novo := []int {}
	for i := 0; i < 202; i ++{
		numeros = append(numeros, i)
		if numeros[i] % 2 != 0{
		novo = append(novo, numeros[i])
	}
	}
	fmt.Printf("%v", novo)
}