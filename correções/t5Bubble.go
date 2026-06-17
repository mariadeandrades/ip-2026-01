package main 

import "fmt"

var N int

func sortBubble(A []int) []int{
    for i := range N {
        for j := 0; j < (N-1-i); j++ {
            if A[j] > A[j+1]{
                A[j], A[j+1] = A[j+1], A[j]
            }
        }
    }
    return A
}

func main() {
	fmt.Scan(&N)

    A := make([]int, N)
    
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
    
    sorted := sortBubble(A)
    
	for i, val := range sorted {
		if i > 0 {
			fmt.Print(" ") 
		}
		fmt.Print(val)
	}
	fmt.Println() 
}