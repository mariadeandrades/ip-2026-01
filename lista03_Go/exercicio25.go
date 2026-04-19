package main

import "fmt"

func main() {
	soma := 0.0
	numerador := 1  //2 elevado a zero
	dezinho := 15  
	for i := 0; i <= 14; i ++{
		denominador := dezinho*dezinho
		termo := numerador/ denominador //termo atual, ele é atualizado a cada iteração
		if i%2 == 0{
			soma += float64(termo)
		}else{
			soma -= float64(termo)
		}
		numerador *= 2
		dezinho = dezinho - 1
	}
	fmt.Printf("O resultado da soma é %.2f.", soma)
}