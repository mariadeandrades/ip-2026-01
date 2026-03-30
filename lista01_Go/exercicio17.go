package main

import "fmt"

func main() {
	var n1, n2, i int
	fmt.Println("Digite dois números, sendo o primeiro par.")
	fmt.Scan(&n1, &n2)
	for n1%2 != 0{
		fmt.Printf("O %d não é par; digite um número par para prosseguir.", n1)
		fmt.Scan(&n1)
	}
	for i = n1; i <= n2; i+=2 {
	fmt.Printf("%d ", i)
	}
}