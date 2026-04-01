package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número.")
	fmt.Scan(&n)
	if n%5 == 0{
		fmt.Printf("O número %v é divisível por 5.", n)
	}else{
		fmt.Printf("O número %v não é divisível por 5.", n)
	}
}