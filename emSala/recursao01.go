//aulas da semana do dia 12/04 ao dia 18/04
/*A LÓGICA DE UMA POTENCIAÇÃO SE COMPORTA DE MANEIRAS DIFERENTES DEPENDENDO DO MÉTODO DE RESOLUÇÃO:
RECURSIVAMENTE: X^N = X * X^(N-1)
USANDO UM LOOP FOR: 

var potencia int = 1

func pot(X, N int) int{
	for i := 1; i <= N; i ++{
		potencia *= X
	}
	return potencia
}
*/ 
package main

import "fmt"

func pot(X, N int) int{
	if N == 0{
		return 1
	}
	return X*pot(X, N-1)
}

func main() {
	var N, X int
	fmt.Println("--ESTE PROGRAMA CALCULA O RESULTADO DE UM NÚMERO X ELEVADO A UM NÚMERO N--")
	fmt.Println("Digite a base (X) que deseja para o cálculo da sua potência:")
	fmt.Println("--OBS: escolha somente valores positivos para a base")
	fmt.Scan(&X)
	for X < 0{
		fmt.Println("Digite somente valores positivos para a base:")
		fmt.Scan(&X)
	}
	fmt.Println("Digite o expoente (N) que deseja para o cálculo da sua potência:")
	fmt.Println("--OBS: escolha somente valores positivos para o expóente!")
	fmt.Scan(&N)
	for N < 0{
		fmt.Println("Digite somente valores positivos para o expoente:")
		fmt.Scan(&N)
	}
	resultado := pot(X, N)
	fmt.Printf("%v elevado a %v resulta em %v.", X, N, resultado)
}