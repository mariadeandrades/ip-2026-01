package main 

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var valor int 
	fmt.Println("Este programa fornece um orçamento para a compra de carros!")
	fmt.Println("Insira o código do feature de design (feature1, feature2, feature3, feature4) que deseja adicionar, de acordo com a seguinte tabela:")
	fmt.Println("feature1: ar condicionado")
	fmt.Println("feature2: pintura metálica")
	fmt.Println("feature3: vidro elétrico")
	fmt.Println("feature4: direção hidráulica")

	//receiving and adjusting strings according to user's imput
	reader := bufio.NewReader(os.Stdin)
    entrada, _ := reader.ReadString('\n')
    entrada = strings.TrimSpace(entrada)

	partes := strings.Fields(entrada)
	for _, opcao := range partes{
		switch opcao {
		case "feature1":
			valor += 1750 
		case "feature2":
			valor += 800
		case "feature3":
			valor += 1200
		case "feature4":
			valor += 2000
		default: 
			fmt.Printf("Opção inválida: %q\n", opcao)
		}
	}
	fmt.Printf("O valor do seu orçamento inicial é de R$%v.", valor)
}