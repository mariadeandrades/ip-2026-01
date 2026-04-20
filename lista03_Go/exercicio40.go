package main

import "fmt"

func main() {
	fmt.Println("-- TABELA DE LUCRO ESPERADO --")
	fmt.Printf("%-15s %-20s %-15s\n", "Preço (R$)", "Ingressos Vendidos", "Lucro (R$)")
	fmt.Println("------------------------------------------------------------")
	const despesaFixa = 300.0
	preco := 6.0
	quantidade := 130
	var lucroMaximo float64
	var precoIdeal float64
	var qtdIdeal int
	for preco >= 1.0 {
		receita := preco * float64(quantidade)
		lucro := receita - despesaFixa
		fmt.Printf("R$ %-12.2f %-20d R$ %-15.2f\n", preco, quantidade, lucro)
		if preco == 6.0 || lucro > lucroMaximo{
			lucroMaximo = lucro
			precoIdeal = preco
			qtdIdeal = quantidade
		}
		preco -= 0.60
		quantidade += 30
	}
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("LUCRO MÁXIMO ESPERADO: R$ %.2f\n", lucroMaximo)
	fmt.Printf("PREÇO IDEAL: R$ %.2f\n", precoIdeal)
	fmt.Printf("INGRESSOS VENDIDOS: %d\n", qtdIdeal)
}