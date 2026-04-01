package main

import "fmt"

func main() {
	var n1, n2, soma, somap, somab int
	fmt.Println("Digite dois números.")
	fmt.Scan(&n1, &n2)
	soma = (n1 + n2)
	somap = (soma + 8)
	somab = (soma - 5)
	if soma > 20{
		fmt.Printf("A soma final é %v", somap)
	}else if soma < 20{
		fmt.Printf("A soma final é %v", soma)
	}else{
		fmt.Printf("A soma final é %v", somab)
	}
}