/*n primo é um n q só é dividido por 1 e por ele mesmo
n/n = 1
n/1 = 1
if a%b = 0{
	a == b
	a é primo
}
*/
/*for i := 0; i < a; i ++{
	for j := 1; j < (a-1); j ++{
		if a%j != 0{
			fmt.Println("Primo")
			break
		}
		fmt.Print("Não é primo")	
	}
}
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número natural que servirá como limite para o intervalo de verificação de números primos:")
	fmt.Scan(&n)
	if n < 2{
		fmt.Printf("%v não é primo.", n)
	}
    totaln := 0
	for i := 2; i < n; i ++{
		ePrimo := true 

		for j := 2; j*j <= i; j ++{
			if i%j == 0{
				ePrimo = false 
				break
			}
		}
		if ePrimo{
			totaln ++
		}
	}
	fmt.Printf("Há %v números primos entre 1 e %v.", totaln, n)
}