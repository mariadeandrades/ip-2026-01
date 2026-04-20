package main

import "fmt"

func main() {
	var graos uint64 = 0
	var graosCasa uint64 = 1

	for quadro := 1; quadro <= 64; quadro++{
		graos += graosCasa
		if quadro < 64{
			graosCasa *= 2
		}
	}
	fmt.Printf("Total final de grãos: %d\n", graos)
	resultado := math.Pow(2, 64) - 1
	fmt.Printf("Em notação científica: %.2e grãos\n", resultado)
}