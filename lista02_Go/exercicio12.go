package main 

import "fmt"

func main() {
	var idade int32 
	fmt.Println("Digite a idade do indivíduo.")
	fmt.Scan(&idade)
	if idade < 3{
		fmt.Println("O indivíduo é classificado como recém-nascido.")
	}else if idade >= 3 && idade < 12{
		fmt.Println("O indivíduo é classificado como criança.")
	}else if idade >= 12 && idade < 20{
		fmt.Println("O indivíduo pe classificado como adolescente.")
	}else if idade >= 20 && idade < 56{
		fmt.Println("O indivíduo é classificado como adulto.")
	}else{
		fmt.Println("O indivíduo é classificado como idoso.")
	}

}