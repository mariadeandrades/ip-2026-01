package main 

import "fmt"

func main () {

	var d int

	fmt.Println("Qual a sua idade?")
	fmt.Scan(&d)
	
	for d < 0{
		fmt.Println("Digite apenas idades não-negativas.")
		fmt.Scan(&d)
	}

	if d < 16{
		fmt.Println("Você se classifica como não-eleitor.")
	}else if d >= 16 && d < 18{
		fmt.Println("Você se classifica como eleitor facultativo.")
	}else if d >= 18 && d < 65{
		fmt.Println("Você se classifica como eleitor obrigatório.")
	}else if d >= 65{
		fmt.Println("Você se classifica como eleitor facultativo.")
	}else{
		fmt.Println("Error")
	}
}