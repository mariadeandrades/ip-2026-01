package main 

import "fmt"

var tipo, comercial, residencial, empresarial string
var conta, consumo float32

func main() {
	fmt.Println("Qual tipo de conta é a sua?")
	fmt.Scan(&tipo) 
	for tipo != "residencial" && tipo != "comercial" && tipo != "empresarial"{
		fmt.Println("Informe um tipo válido de conta.")
	}  
	fmt.Println("Qual o seu consumo mensal de água, em litros?")
	fmt.Scan(&consumo)
	for consumo < 0{
		fmt.Println("Informe um consumo válido.")
	}
	if tipo == "residencial"{
		conta := (5 + 0.5*consumo)
		fmt.Printf("A conta residencial é de %v\n", conta)
	}else if tipo == "comercial"{
		conta := (500 + (consumo - 80)*0.25)
		fmt.Printf("A conta comercial é de %v\n", conta)
	}else if tipo == "empresarial"{
		conta := (800 + (consumo - 100)*0.04)
		fmt.Printf("A conta empresarial é de %v\n", conta)
	}
		
}