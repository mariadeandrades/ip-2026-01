package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Println("-- CÁLCULO DE SENO (SÉRIE DE MACLAURIN) --")
	fmt.Print("Digite o ângulo em radianos: ")
	fmt.Scan(&x)
	for x < 0.0{
		fmt.Println("Informe somente ângulos maiores ou iguais a zero:")
		fmt.Scan(&x)
	}
	seno := 0.0
	termos := 4    
	for i := 0; i < termos; i ++{
		expoente := 2*i + 1
		fatorial := 1.0
		for j := 1; j <= expoente; j ++{
			fatorial *= float64(j)
		}
		potencia := math.Pow(x, float64(expoente))
		termo := potencia/fatorial
		if i%2 == 0{  //q nem a alternância de sinais do ex. 23
			seno += termo
			fmt.Printf("Termo %d: + (%.2f^%d) / %d!\n", i+1, x, expoente, expoente)
		}else{  //o print dentro do for me ajuda a verificar se está funcionando e me ajuda a estudar depois
			seno -= termo
			fmt.Printf("Termo %d: - (%.2f^%d) / %d!\n", i+1, x, expoente, expoente)
		}
	}
	fmt.Printf("Resultado = %.4f\n", seno)
	fmt.Printf("Comparação com math.Sin: %.4f\n", math.Sin(x)) //isso serve pra validar se eu acertei o cálculo ou não
}
