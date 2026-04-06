package main

import "fmt"

var F, C float32

func main() {
	fmt.Println("Qual o valor em Fahrenheit a ser convertido?")
	fmt.Scan(&F)
	C = (5*(F-32)/9)
	fmt.Println(F ,"Fahrenheit equivalem a", C ,"Celsius.")
}