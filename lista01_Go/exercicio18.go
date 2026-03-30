package main

import "fmt"

func main() {
	var a1, r, n int
	fmt.Println("Qual o primeiro termo da P.A., sua razão e seu núemero de termos?")
	fmt.Scan(&a1, &r, &n)
	for i := 0; i < n; i++ {
		term := a1 + i*r
		fmt.Printf("%d ", term)
	}
}