package main

import "fmt"

var valores[numvalores] int
const numvalores int = 20
var soma int = 0
var i int

func main() {
	for i = 1; i <= 20; i ++{
		soma += i
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n%v", soma)
}