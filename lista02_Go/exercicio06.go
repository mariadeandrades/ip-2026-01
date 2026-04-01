package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Digite dois números.")
	fmt.Scan(&a, &b)
	for b == 0{
		fmt.Println("O número b não pode ser nulo, pois não há divisão por zero.\nDigite outro número.")
		fmt.Scan(&b)
	}
	if a%b == 0{
		fmt.Printf("%v é divisível por %v.", a, b)
	}else{
		fmt.Printf("%v não é divísel por %v.", a, b)
	}
}