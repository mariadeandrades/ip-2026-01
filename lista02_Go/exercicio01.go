package main 

import "fmt"

func main() {
	var n1, n2 int
	fmt.Println("Digite dois números inteiros distintos.")
	fmt.Scan(&n1, &n2)
	for n1 == n2{
		fmt.Println("Os números informados são iguais. Digite dois números diferentes entre si.")
		fmt.Scan(&n1, &n2)
	}
	if n1%2 == 0 && n2%2 == 0{
		fmt.Println("Os números informados são pares.")
	}else if n1%2 == 0 && n1%2 != 0{
		fmt.Println("O primeiro número é par e o segundo não é par.")
	}else if n1%2 != 0 && n2%2 == 0{
		fmt.Println("O primeiro número não é par e o segundo é par")
	}else{
		fmt.Println("Nenhum dos números informados é par.")
	}
}