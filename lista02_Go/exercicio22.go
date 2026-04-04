package main

import "fmt"

const min, he float64 = 1788, 10
var m int
var sh, sb, sl, h float64

func main() {
	fmt.Println("Digite seu número de matrícula de funcionário, com três dígitos,\ne a quantidade de horas-extras trabalhadas:")
	fmt.Scan(&m, &h)
	sh = (he*h)
	sb = (3*min + sh)
	if sb > 1500 && sb <= 200{
		sl = (sb - (0.12*sb))
	}else if sb > 200{
		sl = (sb - (0.12*sb) - (0.2*sb))
	}else{
		sl = sb
	}
	fmt.Printf("O seu salário bruto é %.2f e seu salário líquido é %.2f.", sb, sl)
}