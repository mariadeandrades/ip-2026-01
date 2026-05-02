package main

import (
	"fmt"
	"slices"
)

const prim int = 10
const sec int = 5

func main() {
	arrayA := make([]int, prim)
	arrayB := make([]int, sec)
	paresA := []int {}
	imparesA := []int {}
	fmt.Println("Digite, a seguir, 10 números para compor o primeiro array:")
	for i := 0; i < prim; i ++{
	fmt.Println("Digite o", i + 1,"° número.")
	fmt.Scan(&arrayA[i])
	if arrayA[i] % 2 == 0{
		paresA = append(paresA, arrayA[i])
	}else{
		imparesA = append(imparesA, arrayA[i])
	}
	}
	fmt.Println("Digite, a seguir, 5 números para compor o segundo array:")
	for j := 0; j < sec; j ++{
	fmt.Println("Digite o", j + 1,"° número.")
	fmt.Scan(&arrayB[j])
  }
  result1 := slices.Concat(arrayB[:], paresA)
	result2 := slices.Concat(arrayB[:], imparesA)
	fmt.Printf("Os arrays são:\narray A: %v\narray B: %v\narray (B + ímpares de A): %v\narray (B + pares de A): %v", arrayA, arrayB, result2, result1)
}