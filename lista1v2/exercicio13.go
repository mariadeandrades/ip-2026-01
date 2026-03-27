package main 

import "fmt"

var nota float64

func main() {
	fmt.Println("Digite a nota do aluno, entre 0 e 9.")
	fmt.Scan(&nota)
	for nota < 0 || nota > 9{
		fmt.Println("Digite uma nota válida, entre 0 e 9.")
	}
	if nota >= 9 && nota <= 10{
    fmt.Println("A nota do aluno é de conceito A.")
	}else if nota >= 7.5 && nota <= 8.9{
    fmt.Println("A nota do aluno é de conceito B.")
	}else if nota >= 6 && nota <= 7.4{
    fmt.Println("A nota do aluno é de conceito C.")
	}else{
    fmt.Println("A nota do aluno é de conceito D.")
	}
}