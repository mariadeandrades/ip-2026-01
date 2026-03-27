/*
algoritmo_celsius_fahrenheit

declare F, C float

escreva ("Qual o valor em Fahrenheit a ser convertido?")
receba F

funcao C(F)
    retorne(5(F-32)/9)

escreva (F "FAHRENHEIT EQUIAVEL A" C "CELSIUS.")

fim_algoritmo_celsius_fahrenheit
*/

package main

import "fmt"

var F, C float32

func main() {
	fmt.Println("Qual o valor em Fahrenheit a ser convertido?")
	fmt.Scan(&F)
	C = (5*(F-32)/9)
	fmt.Println(F ,"Fahrenheit equivalem a", C ,"Celsius.")
}