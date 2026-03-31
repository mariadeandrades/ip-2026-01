package main 

import "fmt"

var valores[numValores] int
const numValores int = 5
var soma int = 0

func main() {
	for i := 0; i < numValores; i ++{
	fmt.Println("Digite cinco números, um após o outro.")
	fmt.Scan(&valores[i])
	soma += valores[i] 	
	}
	fmt.Printf("A soma dos números é %v", soma)
}