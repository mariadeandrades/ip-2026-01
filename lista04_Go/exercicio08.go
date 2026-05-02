package main

import (
	"fmt"
	"math"
)

func main() {
	array := make([]int, 15)
	fmt.Println("Digite a seguir 15 números inteiros:")
	novo := []int {}
	for i := 0; i < 15; i ++{
		fmt.Println("Digite o", i + 1 ,"° número:")
		fmt.Scan(&array[i])
		if array[i] > 0{
			raiz := math.Sqrt(float64(array[i]))
			novo = append(novo, int(raiz))
		}else{
			raizinha := -1
			novo = append(novo, raizinha)
		}
	}
	fmt.Printf("%v", novo)
}