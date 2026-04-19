package main

import "fmt" 

func main() {
	var b, e int
	
	fmt.Println("Esta função calcula potências depois de receber uma base e um expoente!")
	fmt.Println("Digite a base desejada:")
	fmt.Scan(&b)
	for b < 2{
		fmt.Println("A base escolhida deve ser maior que 1:")
		fmt.Scan(&b)
	}
	fmt.Println("Digite o expoente desejado:")
	fmt.Scan(&e)
	for e < 2{
		fmt.Println("O expoente deve ser maior que 1:")
		fmt.Scan(&e)
	}
	resultado := 1
	for i := 0; i < e; i ++{
		resultado *= b
	}
	fmt.Printf("%v elevado a %v resulta em %v.", b, e, resultado)
}