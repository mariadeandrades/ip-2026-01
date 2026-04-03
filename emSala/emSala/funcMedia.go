// aula do dia 01/04/2026
package main 

import "fmt"

func triple(n1, n2, n3 int) int {
	return ((n1 + n2 + n3)/ 3)
}

func main() {
	var n1, n2, n3, media int
	fmt.Println("Digite três núemros.")
	fmt.Scan(&n1, &n2, &n3)
	media = triple(n1, n2, n3)
	fmt.Printf("A média dos valores %v, %v e %v é %v.", n1, n2, n3, media)
}