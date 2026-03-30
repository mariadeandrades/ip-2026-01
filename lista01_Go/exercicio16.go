package main 

import "fmt"

func main() {
	var salario, novosalario float32

	fmt.Println("Qual o salário do funcionário em questão?")
	fmt.Scan(&salario)
	for salario < 0{
		fmt.Println("Digite um salário válido, ou seja, não negativo.")
		fmt.Scan(&salario)
	}
	if salario <= 300 {
		novosalario = 1.5*salario
	}else{
		novosalario = 1.3*salario
	}
	fmt.Printf("O novo salário é %f.", novosalario)
}
