package main 

import "fmt"

var salariomin, gastomen, conta, desconto float32

func main() {
	fmt.Println("Qual o valor atual do salário mínimo?")
	fmt.Scan(&salariomin)
	fmt.Println("Qual o gasto de energia mensal da residência, em kW?")
	fmt.Scan(&gastomen)
	conta := (gastomen*1.1347)
	desconto := (conta*0.9)
	fmt.Println("O salário mínimo atualmente é de", salariomin)
	fmt.Println("A conta após a aplicação do desconto é de", desconto)
}