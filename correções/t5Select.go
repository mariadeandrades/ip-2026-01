package main

import "fmt"

var N int

func sortSelect(A []int) []int{
    for i := 0; i < N-1; i++ {
        min := i

        for j := i+1; j < N; j++ {
            if A[j] < A[min]{
                min = j
            }
            
        }
		A[i], A[min] = A[min], A[i]
    }
    return A
}

func main() {
	fmt.Scan(&N)
    A := make([]int, N)

    for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
}
    result := sortSelect(A)
    for i, val := range result {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}
