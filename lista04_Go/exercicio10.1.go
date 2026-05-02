package main

import "fmt"

func main() {
	ant1 := 1
	ant2 := 0
	array := make([]int, 50)
	array[0] = 0
	array[1] = 1
	for i := 2; i < 50; i++ {
		atual := ant1 + ant2
		array[i] = atual
		ant2, ant1 = ant1, atual
	}
	fmt.Printf("%v", array)
}