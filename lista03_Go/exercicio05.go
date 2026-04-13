package main

import "fmt"

type Pessoa struct{
	nome string
	idade int
	altura float64
	peso float64
}

var pessoa Pessoa 
var nome string
var altura float64
var peso float64
var idade int

func main() {
	var somah float64 = 0
	var lista []Pessoa

	lista = []Pessoa {}
	 for{
		fmt.Println("Qual o seu nome?")
		fmt.Scan(&nome)
		if nome == "FIM" || nome == "fim"{
			break
		}
	pessoa.nome = nome
	fmt.Println("Qual a sua altura?")
	fmt.Scan(&altura)
	pessoa.altura = altura 
	fmt.Println("Qual o seu peso?")
	fmt.Scan(&peso)
	pessoa.peso = peso
	fmt.Println("Qual a sua idade?")
	fmt.Scan(&idade)
	pessoa.idade = idade
	lista = append(lista, pessoa)
	}
	for _, valor := range(lista) {
	fmt.Printf("Nome: %s\nAltura: %v metros\nPeso: %v Kg\nIdade: %v anos\n\n", valor.nome, valor.altura, valor.peso, valor.idade)
	}
	for _, pessoa := range(lista){
		if pessoa.idade >= 10 && pessoa.idade <= 20{
			somah += pessoa.altura
		}
	}
	fmt.Printf("A soma das alturas das pessoas entre 10 e 20 anos é %.2f.\n", somah)

	listasix := [] Pessoa {}
	for pessoa.idade >= 50{
		listasix = append(listasix, pessoa)
	}
	fmt.Println("Há", len(listasix), "pessoas com mais de cinquenta anos cadastradas.")
}


/*não sei oq fazer agora:
preciso reunir todas as pessoas com mais de 50 anos em uma lista
preciso calcular a média das alturas das pessoas entre 10 e 20 anos --> feito
preciso calcular a porcentagem de pessoas que pesam menos de 40 kg
*/
