package main 

import "fmt"
 
var n1, n2, n3, n4 float32

func main() {
	fmt.Println("Quais são os quatro valores que compõe a matriz 2x2?\nDigite-os na seguinte ordem: primeiro os dois números da primeira linha separados por espaço,\nem seguida os dois números da segunda linha separados por espaço.")
	fmt.Scan(&n1, &n2, &n3, &n4)
	Det := ((n1*n4)-(n2*n3))
	fmt.Println("O determinante da matriz é", Det)
}