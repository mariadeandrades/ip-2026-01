package main

import "fmt"

func main() {
	var salarioCarlos float64

	fmt.Println("Digite o salário de Carlos:")
	fmt.Scan(&salarioCarlos)
	salarioJoao := salarioCarlos / 3

	valorCarlos := salarioCarlos //os capitais iniciais são os salários deles
	valorJoao := salarioJoao
	meses := 0
	for valorJoao < valorCarlos {
		valorCarlos *= 1.02 
		valorJoao *= 1.05   
		meses++
	}

	fmt.Printf("São necessários %d meses para que João iguale ou ultrapasse Carlos.\n", meses)
	fmt.Printf("Valor de Carlos após %d meses: R$ %.2f\n", meses, valorCarlos)
	fmt.Printf("Valor de João após %d meses: R$ %.2f\n", meses, valorJoao)
}