package main 

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Digite três números.")
	fmt.Scan(&a, &b, &c)
	for a == b && b == c{
		fmt.Println("Os números informados são iguais. Digite três números diferentes entre si.")
		fmt.Scan(&a, &b, &c)
	}
	if a > b && b > c{
		maior := a
		inter := b 
		menor := c 
		fmt.Printf("A ordem dos valores é %v, %v, %v.", menor, inter, maior)
	}else if a > c && c > b{
		maior := a
		inter := c 
		menor := b 
		fmt.Printf("A ordem dos valores é %v, %v, %v.", menor, inter, maior)
	}else if b > a && a > c{
		maior := b
		inter := a 
		menor := c 
		fmt.Printf("A ordem dos valores é %v, %v, %v.", menor, inter, maior)
	}else if b > c && c > a{
		maior := b
		inter := c
		menor := a
		fmt.Printf("A ordem dos valores é %v, %v, %v.", menor, inter, maior)
	}else if c > a && a > b{
		maior := c
		inter := a 
		menor := b 
		fmt.Printf("A ordem dos valores é %v, %v, %v.", menor, inter, maior)
	}else{
		maior := c 
		inter := b 
		menor := a 
		fmt.Printf("A ordem dos valores é %v, %v, %v.", menor, inter, maior)
	}
}