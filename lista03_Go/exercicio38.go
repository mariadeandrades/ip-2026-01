package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g, h, i, j, k int
	var d1, d2 int
	fmt.Println("--ESTE PROGRAMA VERIFICA A VALIDADE DE CPFs--")
	fmt.Println("Digite os 11 dígitos do CPF a ser verificado, separados por espaços:")
	fmt.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i, &j, &k)
	soma := (2*i + 3*h + 4*g + 5*f + 6*e + 7*d + 8*c+ 9*b + 10*a)
	resto := soma%11
	if resto < 2{
		d1 = 0
	}else{
		d1 = (11 - resto)
	}
	soma2 := (2*d1 + 3*i + 4*h + 5*g + 6*f + 7*e + 8*d + 9*c + 10*b + 11*a)
	resto2 := soma2%11
	d2 = (11 - resto2)
	if d1 == j && d2 == k{
		fmt.Printf("O CPF %v%v%v%v%v%v%v%v%v%v%v é VÁLIDO.", a, b, c, d, e, f, g, h, i, j, k)
	}else{
		fmt.Printf("O CPF %v%v%v%v%v%v%v%v%v%v%v é INVÁLIDO.", a, b, c, d, e, f, g, h, i, j, k)
	}
}