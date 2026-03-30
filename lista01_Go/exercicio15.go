package main 

import "fmt"

var x, n, i int

func main() {
	fmt.Println("Qual é o número?")
    fmt.Scan(&n)
    for i := 0; i <= n; i ++ {
        fmt.Printf("%d ", i)
    }
    if n%2 == 0{
        x := n*n
        fmt.Printf("\nO quadrado do número %d é %d.\n", n, x)
    }else{
        fmt.Printf("O número %d náo é par.", n)
    }
}