package main

import "fmt"

func main() {
	var matriz = [10][10]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		{3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
		{4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
		{5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
		{6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
		{7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
		{8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
	}
	slice := []int {}
	for i := 0; i < len(matriz); i++{
        for j := 0; j < len(matriz[i]); j++{
            slice = append(slice, matriz[i][j])
        }
    }
    fmt.Printf("Slice resultante: %v\n", slice) //assim fica desorganizado, tentar deixar o print mais amigável
	fmt.Println("Slice Organizado:")
	for i, v := range slice {
    	fmt.Printf("[%02d]: %d  ", i, v)  
    	if i % 10 == 9 {
        	fmt.Println() 
		}
	}
}