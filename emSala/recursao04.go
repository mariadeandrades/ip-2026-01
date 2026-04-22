package main

import "fmt"

func convertBin(n int) string{
	if n == 0{
		return "0"
	}
	if n == 1{
		return "1"
	}
	return convertBin(n/2) + fmt.Sprintf("%d",n%2) //números binários podem ser calculados pelas sucessivas divisões de um número por 2, lemos seus restos de trás pra frente
}

func main() {
	var n int
	fmt.Println("Digite um valor inteiro positivo, que será convertido para o seu equivalente binário:")
	fmt.Scan(&n)
	for n < 0{
		fmt.Println("Digite somente valores positivos:")
		fmt.Scan(&n)
	}
	resultado := convertBin(n)
	fmt.Printf("O equivalente binário de %d é: %s\n", n, resultado)
}