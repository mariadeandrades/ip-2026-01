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
	lista := []Pessoa {}
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
	
}

//não sei oq fazer agora