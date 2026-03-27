package main

import "fmt"

var n int32

func main() {
	fmt.Println("Digite o número que vai passar por verificação de divisibilidade.")
	fmt.Scan(&n)
	if n % 3 == 0 && n % 5 == 0{
		fmt.Printf("O número %d é divisível por 3 e por 5.", n) 
	}else if n % 3 == 0{
		fmt.Printf("O número %d é divísvel por 3.", n)
	}else if n % 5 == 0{
		fmt.Printf("O número %d é divisível por 5.", n)
	}else{
		fmt.Printf("O número %d não é divisível por 3 nem por 5.", n)
	}
}