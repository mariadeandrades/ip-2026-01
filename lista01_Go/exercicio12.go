package main

import "fmt"

var horas float64
var conta float64

func main() {
	fmt.Println("Qual foi o seu tempo de uso?")
	fmt.Scan(&horas)
	if horas < 3{
		conta = 5*(horas)
	}else{
		conta = 10*(horas/3)
    }
	fmt.Printf("O valor da conta é %.2f.", conta)
}