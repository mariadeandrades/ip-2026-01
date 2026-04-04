package main 

import "fmt"

var m int
var n1, n2, n3, me, mf float64
var conceito string

func main() {
	fmt.Println("Digite a seguir seu número de matrícula com 5 dígitos:")
	fmt.Scan(&m)
	fmt.Println("Digite suas notas 1, 2 e 3, bem como a média obtida com a resolução\ndos exercícios, nesta ordem:")
	fmt.Scan(&n1, &n2, &n3, &me)
	mf = ((n1 + n2*2 + n3*3 + me)/ 7)

	if mf <= 10 && mf >= 9{
		conceito = "A"
	}else if mf < 9 && mf >= 7.5{
		conceito = "B"
	}else if mf < 7.5 && mf >= 6{
		conceito = "C"
	}else if mf < 6 && mf >= 4{
		conceito = "D"
	}else{
		conceito = "E"
	}

	switch conceito{
	case "A":
		fmt.Printf("Aluno %v, sua média final é de conceito %s e valor %.2f.\n", m, conceito, mf)
		fmt.Println("Você está aprovado!")
	case "B":
		fmt.Printf("Aluno %v, sua média final é de conceito %s e valor %.2f.\n", m, conceito, mf)
		fmt.Println("Você está aprovado!")
	case "C":
		fmt.Printf("Aluno %v, sua média final é de conceito %s e valor %.2f.\n", m, conceito, mf)
		fmt.Println("Você está aprovado!")
	case "D":
		fmt.Printf("Aluno %v, sua média final é de conceito %s e valor %.2f.\n", m, conceito, mf)
		fmt.Println("Você está reprovado.")
	case "E":
		fmt.Printf("Aluno %v, sua média final é de conceito %s e valor %.2f.\n", m, conceito, mf)
		fmt.Println("Você está reprovado.")
	default:
		fmt.Println("Error")
	}
}