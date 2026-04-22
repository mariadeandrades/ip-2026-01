package main

import "fmt"

func inverter(arr[] int, inicio int, fim int) int{
	if inicio >= fim{
		return 0
	}
	arr [inicio], arr [fim] = arr [fim], arr[inicio]
	return 1 + inverter(arr, inicio+1, fim-1)
}

func main() {
	array := []int {1, 2, 3, 4, 5}
	new := inverter(array, 0, len(array)-1)
	fmt.Printf("O array invertido é %v.", array)
	fmt.Printf("\nPara realizar a inversão, foram feiitas %v trocas.", new)
}