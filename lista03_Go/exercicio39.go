package main

import "fmt"

func main() {
	fmt.Println("-- SISTEMA DE CONTROLE DE PESO DO FRIGORÍFICO --")

	var id, idGordo, idMagro int
	var peso, pesoGordo, pesoMagro float64

	const totalBois = 90

	for i := 1; i <= totalBois; i++ {
		fmt.Printf("Dados do boi %d:\n", i)
		fmt.Print("Número de identificação: ")
		fmt.Scan(&id)
		fmt.Print("Peso (kg): ")
		fmt.Scan(&peso)
		if i == 1{
			idGordo = id
			idMagro = id
			pesoGordo = peso
			pesoMagro = peso
		} else{
			if peso > pesoGordo{
				pesoGordo = peso
				idGordo = id
			}
			if peso < pesoMagro{
				pesoMagro = peso
				idMagro = id
			}
		}
	}

	fmt.Println("--- RESULTADO FINAL ---")
	fmt.Printf("Boi mais GORDO: ID %d com %.2f kg\n", idGordo, pesoGordo)
	fmt.Printf("Boi mais MAGRO: ID %d com %.2f kg\n", idMagro, pesoMagro)
}