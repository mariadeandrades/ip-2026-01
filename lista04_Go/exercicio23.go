package main

import "fmt"

func main() {
	const totalPoltronas = 24
	var janela [totalPoltronas]int   
	var corredor [totalPoltronas]int 

	for {
		cheioJanela := estaCheio(janela)
		cheioCorredor := estaCheio(corredor)
		if cheioJanela && cheioCorredor {
			fmt.Println("\nÔNIBUS COMPLETAMENTE LOTADO!")
			break
		}
		fmt.Println("\n--- SISTEMA DE VENDAS DE PASSAGENS ---")
		fmt.Println("1. Comprar na Janela")
		fmt.Println("2. Comprar no Corredor")
		fmt.Println("3. Finalizar Sistema")
		fmt.Print("Escolha uma opção: ")
		var opcao int
		fmt.Scan(&opcao)
		if opcao == 3 {
			break
		}
		switch opcao {
		case 1: 
			if cheioJanela {
				fmt.Println("Desculpe, não há mais poltronas livres na JANELA.")
			} else {
				venderPoltrona(&janela, "JANELA")
			}
		case 2: 
			if cheioCorredor {
				fmt.Println("Desculpe, não há mais poltronas livres no CORREDOR.")
			} else {
				venderPoltrona(&corredor, "CORREDOR")
			}
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
func estaCheio(v [24]int) bool {
	for _, p := range v {
		if p == 0 {
			return false
		}
	}
	return true
}
func venderPoltrona(v *[24]int, tipo string) {
	fmt.Printf("\nPoltronas livres no(a) %s: ", tipo)
	for i, p := range v {
		if p == 0 {
			fmt.Printf("[%d] ", i+1) 
		}
	}
	var escolha int
	fmt.Print("\nQual poltrona deseja comprar? ")
	fmt.Scan(&escolha)
	if escolha < 1 || escolha > 24 {
		fmt.Println("Número de poltrona inválido!")
	} else if v[escolha-1] == 1 {
		fmt.Println("Esta poltrona já está ocupada. Escolha outra.")
	} else {
		v[escolha-1] = 1 
		fmt.Printf("Venda realizada com sucesso! Poltrona %d (%s) reservada.\n", escolha, tipo)
	}
}