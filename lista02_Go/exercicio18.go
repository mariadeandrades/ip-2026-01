// funciona, mas da próxima vez quero adicionar verificações para q o código aceite pequenas
// variações de escrita no input do usuário
package main

import "fmt"

func main() {
	var valor, dia float32
	var tipo string
	const normal float32 = 10
	fmt.Println("Em qual dia da semana deseja alugar o DVD? Utilize a tabela de correspindência a seguir para digitar o dia:\nsegunda-feira : 2\nterça-feira : 3\nquarta-feira : 4\nquinta-feira : 5\nsexta-feira : 6\nsábado : 7\ndomingo : 1")
	fmt.Scan(&dia)
	fmt.Println("Qual tipo de DVD deseja alugar: comum ou lançamento?")
	fmt.Scan(&tipo)
	switch dia{
	case 1:
		if tipo == "comum"{
			valor = normal
		}else if tipo == "lançamento"{
			valor = (normal*1.15)
		}else{
			fmt.Println("Digite apenas tipos válidos de DVDs.")
		}
	case 2:
		if tipo == "comum"{
			valor = (normal*0.6)
		}else if tipo == "lançamento"{
			valor = (normal*0.69)
		}else{
			fmt.Println("Digite apenas tipos válidos de DVDs.")
		}
	case 3:
		if tipo == "comum"{
			valor = (normal*0.6)
		}else if tipo == "lançamento"{
			valor = (normal*0.69)
		}else{
			fmt.Println("Digite apenas tipos válidos de DVDs.")
		}
	case 4:
		if tipo == "comum"{
			valor = normal
		}else if tipo == "lançamento"{
			valor = (normal*1.15)
		}else{
			fmt.Println("Digite apenas tipos válidos de DVDs.")
		}
	case 5:
		if tipo == "comum"{
			valor = (normal*0.60)
		}else if tipo == "lançamento"{
			valor = (normal*0.69)
		}else{
			fmt.Println("Digite apenas tipos válidos de DVDs.")
		}
	case 6:
		if tipo == "comum"{
			valor = normal
		}else if tipo == "lançamento"{
			valor = (normal*1.15)
		}else{
			fmt.Println("Digite apenas tipos válidos de DVDs.")
		}
	case 7:
		if tipo == "comum"{
			valor = normal
		}else if tipo == "lançamento"{
			valor = (normal*1.15)
		}else{
			fmt.Println("Digite apenas tipos válidos de DVDs.")
		}
	}
	fmt.Printf("A sua compra de DVD do tipo %s terá o custo de R$%.2f.", tipo, valor)
}

