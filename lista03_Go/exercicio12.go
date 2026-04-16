//primeiro denominador = 0! = 1
//preciso inverter o sinal a cada iteração
package main

import "fmt"

func main() {
	var x float64

	fmt.Println("Digite um número positivo:")
	fmt.Scan(&x)
	for x < 0{
		fmt.Println("Digite apenas números maiores que zero:")
		fmt.Scan(&x)
	}
	
	soma := 0.0
	fatorial := 1.0
	sinal := 1.0
	for i := 0; i < 20; i ++{  //aqui eu não posso usar i = 1; i < 20 direto pq i = 0 representa 0!, que faz parte da sequência
		if i > 0{
		fatorial *= float64(i)
		}

	termo := (x*sinal/fatorial)
	soma += termo  //usei uma estrutura parecida com um contador; lembrar disso
	sinal *= -1
	}
	fmt.Printf("Soma = %.2f", soma)
}