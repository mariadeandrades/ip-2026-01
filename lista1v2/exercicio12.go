/*
algoritmo_horas_charrete

escreva ("Qual a quantidade de horas completas de uso?")
receba horas

funcao valor(horas)
se horas <= 3:
    entao valor = 5*horas
se horas = 3:
    entao valor = 10
se horas > 3:
    entao valor = (horas % 3)*5 + #não sei como terminar
*/

package main

import "fmt"

var horas float64
var conta int

func main() {
	fmt.Println("Qual foi o seu tempo de uso?")
	fmt.Scan(&horas)
	if horas < 3{
		conta = 5*int(horas)
	}else{
		conta = 10*int(horas/3)
    }
	fmt.Printf("O valor da conta é %d.", conta)
}