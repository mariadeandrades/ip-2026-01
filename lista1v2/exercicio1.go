// algoritmo_aprova_reprova

// declare media, nota1, nota2, nota3 numerico

// escreva ("Quais são suas notas? (Digite-as espaçadamente)")
// leia nota1 nota2 nota3

// funcao media(nota1, nota2, nota3)
//     retorne ((nota1 + nota2 + nota3)/3) "\n"
// se media >= 6:
//   escreva ("APROVADO!\n")
// senão
//   escreva ("REPROVADO!\n")

// fim_algoritmo_aprova_reprova
package main

import "fmt"

func main() {
	var nota1, nota2, nota3 float32
	fmt.Println("Quais são suas notas? (Digite-as espaçadamente: 10 7 6)")
	fmt.Scan(&nota1, &nota2, &nota3)
}
func media(nota1, nota2, nota3 float32) float32 {
	return (nota1 + nota2 + nota3) / 3
}	