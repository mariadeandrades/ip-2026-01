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
var altura, peso, contador1, media, fracao float64
var idade, contador2, contador3, size int

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
	size = len(lista)
	}
	for _, valor := range(lista) {
	fmt.Printf("Nome: %s\nAltura: %v metros\nPeso: %v Kg\nIdade: %v anos\n\n", valor.nome, valor.altura, valor.peso, valor.idade)
	}
	for _, pessoa := range(lista){
		if pessoa.idade >= 10 && pessoa.idade <= 20{
			contador1 += 1
			somah += pessoa.altura
			media = (somah/contador1)
		}
		if pessoa.idade > 50{
			contador2 += 1
		}
		if pessoa.peso < 40{
			contador3 += 1
			fracao = float64((contador3/size))  //transformar o resultado pra float64
		}
	}
	fmt.Printf("A média das alturas das pessoas entre 10 e 20 anos é %.2f.\n", media)
	fmt.Printf("Há %v pessoas com mais de 50 anos no cadastro.", contador2)
	fmt.Printf("A porcentagem de pessoas que pesam menos de 40 Kg no cadastro é de %.2f", fracao)
}
