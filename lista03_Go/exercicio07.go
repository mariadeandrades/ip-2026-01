package main

import "fmt"

var n, min, max, total, contarpares, contarimpares int

func main() {
	var numeros[] int
	var pares[] int
	var impares[] int

	numeros = []int {} 
	fmt.Println("Digite o primeiro número natural qualquer:")
	fmt.Scan(&n)
	numeros = append(numeros, n)
	max = n
	min = n 
	for{
		fmt.Println("Digite o próximo número natural qualquer:")
		fmt.Scan(&n)
		if n > max {
            max = n
        }
        if n < min {
            min = n
        } 
		if n == 30000{
			break
		}
		numeros = append(numeros, n) //deixar o append dentro do loop, caso contrário só o último valor será considerado
	}
	total = len(numeros)
	soma := 0
	for _, num := range numeros{
		soma += num
	}
	media := (soma/total)
	pares = []int {}
	if n %2 == 0{
		pares = append(pares, n)
		contarpares += 1
	}
	somapares := 0
	for _, par := range pares{
		somapares += par
	}
	mediapares := (somapares/contarpares)
	impares = []int {}
	if n %2 != 0{
		impares = append(impares, n)
		contarimpares += 1
	}
	fracimpares := (contarimpares/total)

	fmt.Printf("QUANTIDADE DE VALORES: %v\nSOMA DOS VALORES: %v\nMÉDIA DOS VALORES: %v\nMAIOR VALOR: %v\nMENOR VALOR: %v\nMÉDIA PARES: %v\nFRAÇÃO ÍMPARES: %v", total, soma, media, max, min, mediapares, fracimpares)
}
