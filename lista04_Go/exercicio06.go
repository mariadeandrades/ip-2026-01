package main

import "fmt"

func main() {
	var numeros []int
	for i := 100; i >= 1; i --{
		numeros = append(numeros, i)
	}
	fmt.Printf("%v", numeros)
}