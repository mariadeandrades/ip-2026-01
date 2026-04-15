package main

import "fmt"

type Aluno struct{
	nome string 
	n1 float64
	n2 float64
	media float64
	soma float64
}

var aluno Aluno
var nome string
var n1, n2, media, soma, somanotas float64
var size, contadorrep, contadorex, contadorap int

func main() {
	var lista []Aluno

	lista = []Aluno {}
	for{
		fmt.Println("Qual o nome do aluno?")
		fmt.Scan(&nome)
		if nome == "fim" || nome == "FIM" || nome == "f" || nome == "F"{
			break
		}
	aluno.nome = nome
	fmt.Println("Qual a primeira nota do aluno?")
	fmt.Scan(&n1)
	aluno.n1 = n1
	fmt.Println("Qual a segunda nota do aluno?")
	fmt.Scan(&n2)
	aluno.n2 = n2
	soma := (n1 + n2)
	aluno.soma = soma
	media := (aluno.soma/2)
	aluno.media = media
	lista = append(lista, aluno)
	size = len(lista)
	}
	for _, aluno := range(lista){
		if aluno.media <= 3{
			contadorrep += 1
		}else if aluno.media > 3 && aluno.media < 7{
			contadorex += 1 
		}else if aluno.media >= 7{
			contadorap += 1
		}
	}
	somanotas := 0.0 //garante q o Go entenda q estou declarando e inicializando uma variável float, e não int
	for _, aluno:= range(lista){
		somanotas += aluno.soma
	}
	medianotas := (somanotas/2)

	fmt.Println("---Média Aritmética---\nNota até 3: Reprovado\nNota entre 3 e 7: Exame\nNota de 7 para cima: Aprovado")
	fmt.Printf("ALUNOS APROVADOS: %v\nALUNOS EM EXAME: %v\nALUNOS REPROVADOS: %v", contadorap, contadorex, contadorrep)
	fmt.Printf("\nA média da classe é %.2f.", medianotas)
}


