// aula do dia 01/04/2026
package main 

import "fmt"

func somar(n1, n2 int) int {
	return (n1 + n2)
}

func main() {
	var n1, n2, soma int
	fmt.Println("Digite dois valores.")
	fmt.Scan(&n1, &n2)
	soma = somar(n1, n2)
	fmt.Printf("A soma dos valores %v e %v é %v.", n1, n2, soma)
}
