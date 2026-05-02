package main

import "fmt"

func main() {
	array := make([]int, 10)
	fmt.Println("Digite, a seguir, os 10 números para compor o array:")
	for i := 0; i < 10; i ++{
		fmt.Println("Digite o", i + 1 ,"° número:")
		fmt.Scan(&array[i])
	}
	freq := make(map[int]int)
	for _, num := range array {
		freq[num]++
	}
	fmt.Printf("Array original: %v\n", array)
	fmt.Println("Números repetidos e suas contagens:")
    for num, count := range freq {
        if count > 1{
            fmt.Printf("Número %d aparece %d vezes\n", num, count)
		}
    }
}