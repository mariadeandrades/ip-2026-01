package main 

import "fmt"

func main() {
	var c, pn float64
	var forma int

	fmt.Println("Este programa calcula quanto você pagará pela sua compra de acordo\ncom a forma de pagamento escolhida!")
	fmt.Println("Qual o valor da sua compra?")
	fmt.Scan(&pn)
	fmt.Println("Escolha a forma de pagamento de acordo com a lista numerada a seguir\n digitando o número correspondente:")
	fmt.Println("(1) : À vista em dinheiro ou cheque\n(2) : À vista no cartão de crédito\n(3) : Parcelado em duas vezes\n(4) : Parcelado em três vezes")
	fmt.Scan(&forma)
	switch forma{
	case 1:
		c = (pn*0.9)
	case 2:
		c = (pn*0.95)
	case 3:
		c = pn
	case 4:
		c = (pn*1.1)
	default:
		fmt.Println("Forma de pagamento inválida.")
	}
	fmt.Printf("Com a forma de pagamento escolhida (%v), você pagará R$%.2f pela sua compra.", forma, c)
}