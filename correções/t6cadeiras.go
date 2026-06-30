package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	max := 0
	count := 0

	for i := N - 1; i >= 0; i-- {
		if A[i] < max {
			count++
		} else {
			max = A[i]
		}
	}

	fmt.Println(count)
}