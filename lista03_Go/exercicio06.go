package main

import "fmt"

func consect(num, int) (bool, int, int, bool) {
	if num < 0{
		return false, 0, 0, 0
	}
	for n := 0; n++ {
		produto := (n*(n+1)*(n+2))
	}
}

func main() {
	fmt.Println("Digite um número não-negativo:")
	fmt.Scan(&n)
	if n < 0{
		fmt.Println("O número não é triangular.")
	}
}
/*o professor quer um programa que verifique se um número
qualquwer é formado pela multiplicação de três consecutivos
ex: 6 = 1*2*3
a entrada vai ser o número que vai sofrer verificação, eu não sei como
fazer o programa buscar essa correspondência 
