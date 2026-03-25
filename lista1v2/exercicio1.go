package main

import "fmt"

func media(nota1, nota2, nota3 float64) float64 {
	return (nota1 + nota2 + nota3) / 3
}	

func main() {
	var nota1, nota2, nota3 float64
	var med float64
	fmt.Println("Quais são suas notas? (Digite-as espaçadamente: 10 7 6)")
	fmt.Scan(&nota1, &nota2, &nota3)
	for (nota1 > 10 || nota2 > 10 || nota3 > 10 || nota1 < 0 || nota2 < 0 || nota3 < 0) {
		fmt.Println("Digite valores válidos, ou seja, maiores que zero e menores que dez.")
		fmt.Scan(&nota1, &nota2, &nota3)
	}
	med = media(nota1, nota2, nota3)   
	if med < 6{
		fmt.Println("Você está reprovado.")
	}else{
		fmt.Println("Você está aprovado.")
	}
}