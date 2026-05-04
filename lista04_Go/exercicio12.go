package main

import "fmt"

func main() {
	notas := make([]int, 15)
	fmt.Println("Digite, a seguir, as notas dos 15 alunos:")
	for i := 0; i < 15; i ++{
		fmt.Println("Digite a", i + 1 ,"° nota:")
		fmt.Scan(&notas[i])
		for notas[i] < 0 || notas[i] > 10{
			fmt.Println("Nota inválida, digite um valor entre 0 e 10 (inclusos):")
			fmt.Scan(&notas[i])
		}
	}
	contador0, contador1, contador2, contador3, contador4, contador5, contador6, contador7, contador8, contador9, contador10 := 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	for i := 0; i < len(notas); i ++{
		if notas[i] == 0{
			contador0 += 1
		}else if notas[i] == 1{
			contador1 += 1
		}else if notas[i] == 2{
			contador2 += 1
		}else if notas[i] == 3{
			contador3 += 1
		}else if notas[i] == 4{
			contador4 += 1
		}else if notas[i] == 5{
			contador5 += 1
		}else if notas[i] == 6{
			contador6 += 1
		}else if notas[i] == 7{
			contador7 += 1
		}else if notas[i] == 8{
			contador8 += 1
		}else if notas[i] == 9{
			contador9 += 1
		}else{
			contador10 += 1
		}
	}
	//para as frequencias relativas
	fr0 := float64(contador0)/10
	fr1 := float64(contador1)/10
	fr2 := float64(contador2)/10
	fr3 := float64(contador3)/10
	fr4 := float64(contador4)/10
	fr5 := float64(contador5)/10
	fr6 := float64(contador6)/10
	fr7 := float64(contador7)/10
	fr8 := float64(contador8)/10
	fr9 := float64(contador9)/10
	fr10 := float64(contador10)/10

	fmt.Println("--NOTAS--")
	fmt.Printf("\n-- %v --", notas)
	fmt.Println("\n\n--FREQUÊNCIAS ABSOLUTAS--")
	fmt.Printf("\n-- nota 0 = %v  nota 1 = %v  nota 2 = %v  nota 3 = %v  nota 4 = %v  nota 5 = %v  \nnota 6 = %v  nota 7 = %v  nota 8 = %v  nota 9 = %v  nota 10 = %v --", contador0, contador1, contador2, contador3, contador4, contador5, contador6, contador7, contador8, contador9, contador10)
	fmt.Println("\n\n--FREQUÊNCIAS RELATIVAS--")
	fmt.Printf("\n--nota 0 = %v  nota 1 = %v  nota 2 = %v  nota 3 = %v  nota 4 = %v  nota 5 = %v  nota 6 = %v  nota 7 = %v  nota 8 = %v  nota 9 = %v  nota 10 = %v --", fr0, fr1, fr2, fr3, fr4, fr5, fr6, fr7, fr8, fr9, fr10)
}