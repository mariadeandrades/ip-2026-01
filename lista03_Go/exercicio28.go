package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("--ESTE PROGRAMA CALCULA E IMPRIME UM VALOR APROXIMADO PARA PI--")
	somaPi := 0.0
	numerador := 1.0
	for i := 1; i <= 51; i ++{
		denominador := float64(2*i - 1) //denominadores ímpares
		potencia := math.Pow(denominador, float64(3))
		termo := numerador/potencia
		if i%2 != 0{ 
			somaPi += termo
			fmt.Printf("Termo %d: + 1 / %.0f\n", i, potencia)
		}else{  
			somaPi -= termo
			fmt.Printf("Termo %d: - 1 / %.0f\n", i, potencia)
		}
	}
	produto := somaPi*32
	raiz := math.Cbrt(produto)
	fmt.Printf("A soma, que oferece um valor aproximado para Pi, resulta em: %.2f.", raiz)
}