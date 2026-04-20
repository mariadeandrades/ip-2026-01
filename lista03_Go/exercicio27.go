package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Println("--ESTE PROGRAMA CALCULA O COSSENO DE ÂNGULOS QUAISQUER--")
	fmt.Println("Digite, em radianos, o ângulo cujo cosseno deseja calcular:")
	fmt.Scan(&x)
	for x < 0{
		fmt.Println("Digite somente valores de ângulos positivos:")
		fmt.Scan(&x)
	}
	cosseno := 0.0
	termos := 20
	for i := 0; i < termos; i ++{
		expoente := 2*i
		fatorial := 1.0
		for j := 0; j <= expoente; j ++{
			fatorial *= float64(j)
		}
		potencia := math.Pow(x, float64(expoente))
		termo := potencia/fatorial
		if i%2 != 0{ 
			cosseno += termo
			fmt.Printf("Termo %d: + (%.2f^%d) / %d!\n", i+1, x, expoente, expoente)
		}else{  
			cosseno -= termo
			fmt.Printf("Termo %d: - (%.2f^%d) / %d!\n", i+1, x, expoente, expoente)
		}
	}
	fmt.Printf("Resultado = %.4f\n", cosseno)
	fmt.Printf("Comparação com math.Sin: %.4f\n", math.Cos(x))
}