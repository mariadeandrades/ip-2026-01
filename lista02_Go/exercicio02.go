package main 

import "fmt"

func main() {
	var n int
	fmt.Println("Informe um número.")
	fmt.Scan(&n)
	if n > 0{
		fmt.Printf("O número %d é positivo.", n)
	}else if n < 0{
		fmt.Printf("O número %d é negativo.", n)
	}else{
		fmt.Printf("O número %d é nulo.", n)
	}
}