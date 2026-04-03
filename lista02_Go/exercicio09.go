package main 

import "fmt"

func main() {
	var compra, venda float32
	fmt.Println("Qual o valor de compra da mercadoria?")
	fmt.Scan(&compra)
	for compra <= 0{
		fmt.Println("O valor de compra não pode ser nulo ou negativo.\nDigite um valor de compra válido.")
		fmt.Scan(&compra)
	} 
	if compra < 10{
		venda = (1.7*compra)
		fmt.Printf("O valor de venda da mercadoria é de R$%v", venda)
	}else if compra >= 10 && compra < 30{
		venda = (1.5*compra)
		fmt.Printf("O valor de venda da mercadoria é de R$%v", venda)
	}else if compra >= 30 && compra < 50{
		venda = (1.4*compra)
		fmt.Printf("O valor de venda da mercadoria é de R$%v", venda)
	}else{
		venda = (1.3*compra)
		fmt.Printf("O valor de venda da mercadoria é de R$%v", venda)
	}
}