// aula do dia 30/03/2026
package main

import "fmt"

var nota[numNotas] float64
const numNotas int = 3
var soma float64 = 0
var media float64

func main() {
	for i := 0; i < numNotas; i ++{
		fmt.Println("Digite a", i + 1, "° nota do aluno.")
		fmt.Scan(&nota[i])
		soma += nota[i]
	}
	media = (soma/3)
	fmt.Printf("A média das notas do aluno é %f.", media)
}
