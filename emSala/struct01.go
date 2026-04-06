/*package main

import "fmt"

type Criatura struct {
	Nome string
}

func main() {
	c := Criatura{
		Nome: "Jac and friends",
	}
	fmt.Println(c.Nome)
}
Crie uma struct para armazenar o nome, altura e peso ideal de uma pessoa. Depois crie um slice e
preencha seus elementos com structs do tipo pessoa. Leia os dados com uma repetição até ser digitado um nome igual a “FIM”. 
Calcule o peso ideal com base na fórmula: Peso ideal = 72.7 x altura - 58.0. Imprima o slice

*/
package main 

import "fmt"

type Pessoa struct {
	nome string 
	altura float64
	pesoIdeal float64 
}

var pessoa Pessoa 
var nome string 
var altura float64 
var pesoIdeal float64

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
	pesoIdeal = (72.7*pessoa.altura - 58)
	pessoa.pesoIdeal = pesoIdeal
	lista = append(lista, pessoa)
	}
	for _, valor := range(lista) {
	fmt.Printf("As características da pessoa %s são %.2f metros de altura\ne um peso ideal de %.2f quilogramas.\n", valor.nome, valor.altura, valor.pesoIdeal)
	}
}