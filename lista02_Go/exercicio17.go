package main

import "fmt"

func main() {
	var tipo string
	var c, valor float32
	fmt.Println("Este programa calcula contas de água para a empresa Saneago.")
	fmt.Println("Digite qual tipo de consumidor você é (residencial, comercial ou industrial):")
	fmt.Scan(&tipo)
	fmt.Println("Digite o seu consumo de água mensal, em metros cúbicos.")
	fmt.Scan(&c)
	if tipo == "residencial" {
		valor = (5 + c*0.05)
		fmt.Printf("A sua conta do tipo %s, vinculada a um consumo de %.2f metros cúbicos mensais, é de R$%.2f.\n", tipo, c, valor)
	} else if tipo == "comercial" {
		if c > 80 {
			valor = (500 + (c-80)*0.25)
		} else {
			valor = 500
		}
		fmt.Printf("A sua conta do tipo %s, vinculada a um consumo\nde %.2f metros cúbicos mensais, é de R$%.2f.\n", tipo, c, valor)
	} else if tipo == "industrial" {
		if c > 100 {
			valor = (800 + (c-100)*0.04)
		} else {
			valor = 800
		}
		fmt.Printf("A sua conta do tipo %s, vinculada a um consumo de %.2f metros cúbicos mensais, é de R$%.2f.\n", tipo, c, valor)
	} else {
		fmt.Println("Digite apenas tipos válidos de consumidor.")
	}
}