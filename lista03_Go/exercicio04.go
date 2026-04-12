package main

import "fmt"

func main() {
	var n int
	var soma int = 0

	fmt.Println("Digite um número não-negativo:")
	fmt.Scan(&n)

	for i := 1; i < n; i += 2{  //para um número n quadrado perfeito, que tem raiz quadrada m,
		                        //n será igual à soma dos m primeiros números ímpares
		soma += i
	}

	if soma == n{
		fmt.Println("O número é quadrado perfeito.")
	}else{
		fmt.Println("O número não é quadrado perfeito.")
		}
}

/*aqui eu aprendi a evitar deixar if e for um dentro do outro;
separar sempre que possível!*/