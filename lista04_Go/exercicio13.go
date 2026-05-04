package main

import "fmt"

type Funcionario struct {
	Numero int
	Tempo int
}

var funcionario Funcionario
var numero int
var tempo int

func main() {
	var array [100]Funcionario
	var total int
	fmt.Println("Digite o número do funcionário e seu tempo de trabalho:")
	fmt.Println("Caso queira encerrar, digite 0 e 0.")
	for i := 0; i < 100; i ++{
		fmt.Scan(&numero, &tempo)

		if numero == 0 && tempo == 0{
			break
		}
		array[i] = Funcionario{Numero : numero, Tempo : tempo}
		total ++
	}
	if total < 3{
		fmt.Println("Insira dados de pelo menos 3 funcionários.")
		return
	}
	for i := 0; i < total-1; i++ {
		for j := 0; j < total-i-1; j++ {
			if array[j].Tempo > array[j+1].Tempo {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println("\n--Os 3 Empregados Mais Recentes--")
	for k := 0; k < 3; k++ {
		fmt.Printf("%dº Lugar: Empregado %d (%d meses)\n", k+1, array[k].Numero, array[k].Tempo)
	}
}
