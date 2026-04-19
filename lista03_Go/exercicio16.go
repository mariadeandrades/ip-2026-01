package main

import "fmt" 

func main() {
	var n1, n2, N int

	fmt.Println("Digite o primeiro termo da sequência de Fetuccini:")
	fmt.Scan(&n1)
	fmt.Println("Digite o segundo termo da sequência de Fetuccini:")
	fmt.Scan(&n2)
	fmt.Println("Digite o tamanho que deseja para sua sequência de Fetuccini:")
	fmt.Scan(&N)
	
	numeros := []int {n1, n2} //inicializar um slice com dois termos não significa que ele só poderá guardar dois termos
	
	for i := 2; i < N; i ++{  //o loop começa em 2 pq já tenho os dois primeiros termos (0 e 1)
		var next int
		if (i + 1) % 2 != 0 {
			next = numeros[i-1] + numeros[i-2]
		}else{
			next = numeros[i-1] - numeros[i-2]
		}
		numeros = append(numeros, next)
	}
	fmt.Printf("Sequência de Fettuccine: %v", numeros)
}
