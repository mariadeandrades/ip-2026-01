package main

import "fmt"

func potencia(base, expoente int) int {
	resultado := 1 //fazemos isso pq 1 é o elemento neutro da multiplicação
	for i := 0; i < expoente; i++ {  
		resultado *= base   //resultado = resultado * base        
	}  //o expoente é o número de vezes q esse loop se repetirá    
	return resultado
}

func main() {
	var base, expoente int
	fmt.Println("Digite a base da potência.")
	fmt.Scan(&base)
	fmt.Println("Digite o expoente.")
	fmt.Scan(&expoente)
	for expoente < 0 {
		fmt.Println("Esta calculadora só aceita expoentes positivos. Digite um expoente válido:")
		fmt.Scan(&expoente)
	}
	resultado := potencia(base, expoente)
	fmt.Printf("%d elevado a %d é igual a %d\n", base, expoente, resultado)
}