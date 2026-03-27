package main

import "fmt"

var rendat, ptotal, ppopular, pgeral, parqui, pcadeiras, rp, rg, ra, rc float32 

func main(){
	fmt.Println("Qual o público total do jogo?")
	fmt.Scan(&ptotal)
	fmt.Println("Qual a porcentagem de público nas categorias popular, geral, arquibancada e cadeira, respectivamente?")
	fmt.Scan(&ppopular, &pgeral, &parqui, &pcadeiras)
	rp = (ppopular*ptotal)
	rg = (pgeral*ptotal*5)
	ra = (parqui*ptotal*10)
	rc = (pcadeiras*ptotal*20)
	rendat = (rp + rg + ra + rc)
	fmt.Printf("A renda total é de %v", rendat) 
	fmt.Printf("A renda popular é de %v\n", rp)
	fmt.Printf("A renda geral é de %v\n", rg)
	fmt.Printf("A renda das arquibancadas é de %v\n", ra)
	fmt.Printf("A renda das cadeiras é de %v\n", rc)
}