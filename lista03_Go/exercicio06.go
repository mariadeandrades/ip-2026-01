package main

import "fmt"

var num, n, produto int
var resultado bool

func consect(num, int) bool {
	if num < 0{
		return false
	}
	produto := (n*(n+1)*(n+2))
	if produto == num{
		return true
	}else{
		return false
		}
}

func main() {
	fmt.Println("Digite um número natural, isto é, não-negativo:")
	fmt.Scan(&n)
	resultado := consect(num)
	fmt.Println(resultado)
}
/*o professor quer um programa que verifique se um número
qualquwer é formado pela multiplicação de três consecutivos
ex: 6 = 1*2*3
a entrada vai ser o número que vai sofrer verificação, eu não sei como
fazer o programa buscar essa correspondência 
