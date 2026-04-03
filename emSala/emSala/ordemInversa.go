// aula do dia 30/03/2026
package main

import "fmt"

var array[numValores] int
const numValores int = 10

func main() {
	for i := 0; i < numValores; i ++{
		fmt.Println("Digite dez valores.")
		fmt.Scan(&array[i])
	}
	novoArray := []int{}
	for i := 9; i >= 0; i-- {
		novoArray = append(novoArray, array[i])
	}
	fmt.Println("Os números em ordem decrescente são:", novoArray)
}
