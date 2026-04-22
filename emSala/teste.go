package main

import (
	"fmt"
	"strings"
)

type Vendedor struct{
	nome string
	totalVendido float64
	numeroVendas int
	mediaVendas float64
}

var vendedor Vendedor
var nome string
var numerovendas, contadorv, contador2 int
var totalvendido, mediavendas, vendas float64

func main() {
	lista := []Vendedor {}
	for{
		fmt.Println("Qual o nome do vendedor?")
		fmt.Scan(&nome)
		if nome == "fim" || nome == "FIM" || nome == "F" || nome == "f"{
			break
		}
		vendedor.nome = nome
		contadorv += 1
		fmt.Println("Digite quantas vendas o vendedor realizou:")
		fmt.Scan(&numerovendas)
		vendedor.numeroVendas = numerovendas
		fmt.Println("Digite o valor total vendido pelo vendedor:")
		fmt.Scan(&totalvendido)
		vendedor.totalVendido = totalvendido
		mediaVendas = totalvendido/numerovendas
		vendedor.mediaVendas = mediavendas
		lista = append(lista, vendedor)
	}
	for _, vendedor := range(lista){
		vendas += vendedor.totalVendido
	}
	media := vendas/contadorv
	for _, vendedor := range(lista){
		if mediaVendas < media{
			contador2 += 1
		}
	}
	fmt.Printf("%v = total de vendas da emnpresa", vendas)
	fmt.Printf("%v = vendedores abaixo da média", contador2)
}