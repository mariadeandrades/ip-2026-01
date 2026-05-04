package main

import "fmt"

func main() {
	const totalContas = 10
	var codigos [totalContas]int
	var saldos [totalContas]float64
	fmt.Println("--- Cadastro de Contas Bancárias ---")
	for i := 0; i < totalContas; i++ {
		for {
			var codigoTmp int
			fmt.Printf("Digite o código da %dª conta: ", i+1)
			fmt.Scan(&codigoTmp)
			existe := false
			for j := 0; j < i; j++ {
				if codigos[j] == codigoTmp {
					existe = true
					break
				}
			}
			if existe {
				fmt.Println("Erro: Já existe uma conta com este código. Tente outro.")
			} else {
				codigos[i] = codigoTmp
				fmt.Printf("Digite o saldo inicial da conta %d: ", codigoTmp)
				fmt.Scan(&saldos[i])
				break
			}
		}
	}
	for {
		fmt.Println("\n--- MENU BANCÁRIO ---")
		fmt.Println("1. Efetuar depósito")
		fmt.Println("2. Efetuar saque")
		fmt.Println("3. Consultar ativo bancário")
		fmt.Println("4. Finalizar programa")
		fmt.Print("Escolha uma opção: ")
		var opcao int
		fmt.Scan(&opcao)
		if opcao == 4 {
			fmt.Println("Programa finalizado. Até logo!")
			break
		}
		switch opcao {
		case 1, 2: 
			var contaBusca int
			var valor float64
			fmt.Print("Digite o código da conta: ")
			fmt.Scan(&contaBusca)
			indice := -1
			for i := 0; i < totalContas; i++ {
				if codigos[i] == contaBusca {
					indice = i
					break
				}
			}
			if indice == -1 {
				fmt.Println("Conta não encontrada.")
			} else {
				if opcao == 1 {
					fmt.Print("Valor do depósito: ")
					fmt.Scan(&valor)
					saldos[indice] += valor
					fmt.Printf("Depósito realizado! Novo saldo: R$ %.2f\n", saldos[indice])
				} else {
					fmt.Print("Valor do saque: ")
					fmt.Scan(&valor)
					if valor <= saldos[indice] {
						saldos[indice] -= valor
						fmt.Printf("Saque realizado! Novo saldo: R$ %.2f\n", saldos[indice])
					} else {
						fmt.Println("Saldo insuficiente.")
					}
				}
			}
		case 3:
			var ativo float64
			for _, s := range saldos {
				ativo += s
			}
			fmt.Printf("O ativo bancário total é: R$ %.2f\n", ativo)
		default:
			fmt.Println("Opção inválida!")
		}
	}
}