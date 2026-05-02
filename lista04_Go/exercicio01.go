package main

import "fmt"

const numValores int = 10

func main() {
	numeros := make([]int, numValores)
	novo := []int {}
	fmt.Println("--ESTE PROGRAMA RECEBE 1O NÚMEROS E RETORNA AQUELES QUE FOREM SUPERIORES A 50--")
	for i := 0; i < numValores; i ++{
	fmt.Println("Digite o", i + 1,"° número:")
	fmt.Scan(&numeros[i])
	if numeros[i] > 50{
		novo = append(novo, numeros[i])
	}
  }
  for i, valor := range(numeros){
	fmt.Printf("%v° número da lista: %v.\n", 1 + i, valor)
  }
  if len(novo) > 0{
	fmt.Printf("Os números da sequência %v que são maiores que 50 são %v.", numeros, novo)
}else{
	fmt.Println("Nenhum número maior que 50 foi encontrado.")
}
}