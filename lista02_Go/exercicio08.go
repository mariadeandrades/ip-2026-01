package main 

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número entre 20 e 90.")
	fmt.Scan(&n)
	if n > 20 && n < 90{
		fmt.Printf("O número %v está dentro do intervalo.", n)
	}else{
		fmt.Printf("O número %v não está dentro do intervalo.", n)
	}
}