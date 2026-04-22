// aula do dia 01/04/2026
package main 

import "fmt"

func comp(n1, n2, n3 int) int {
		if n1 > n2 && n1 > n3{
			return n1
		}
		if n2 > n1 && n2 > n3{
			return n2
		}
		if n3 > n1 && n3 > n2{
			return n3
		}
		return n1
}
func main() {
	var m, n1, n2, n3 int
	fmt.Println("Digite três números.")
	fmt.Scan(&n1, &n2, &n3)
	m = comp(n1, n2, n3)
	if n1 == n2 && n2 == n3{
		fmt.Printf("Os três números são iguais, portanto o maior valor é: %v\n", m )
	}else{
		fmt.Printf("O maior número é %v", m)
	}
}
