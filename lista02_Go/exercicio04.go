package main 

import (
	"fmt"
 	"math"
)
func main() {
	var n, quadrado, raiz float64
	fmt.Println("Digite um número.")
	fmt.Scan(&n)
	quadrado = (n*n)
	raiz = math.Sqrt(n)
	if n < 0{
		fmt.Printf("O quadrado de %f é %f.", n, quadrado)
	}else{
		fmt.Printf("A raiz do de %f é %f.", n, raiz)
	}

}