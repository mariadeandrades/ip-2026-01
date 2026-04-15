package main

import "fmt"

var num, n, produto int
var resultado bool

func consect(num int) bool {
	if num < 0{
		return false
	}
	for n := 1; ; n ++{
	produto := (n*(n+1)*(n+2))
		if produto == num{
			return true
		}else{
			return false
		}
    }
}

func main() {
	fmt.Println("Digite um número natural, isto é, não-negativo:")
	fmt.Scan(&num)
	resultado := consect(num)
	if resultado == true{
		fmt.Printf("O número %v é formado pelo produto de três números sucessivos!", num)
	}else{
		fmt.Printf("O número %v não é formado pelo produto de três números sucessivos!", num)
	}
}

