package main

import "fmt"

func fibonacci(n int) []int {
    if n <= 0 {
        return nil
    }
    if n == 1 {
        return []int{1}
    }
    if n == 2 {
        return []int{1, 1}
    }

    seq := fibonacci(n - 1)
    seq = append(seq, seq[len(seq)-1]+seq[len(seq)-2])
    return seq
}

func main() {
    seq := fibonacci(50)
    fmt.Println("Primeiros 50 termos da sequência de Fibonacci:")
    for i, v := range seq {
        fmt.Printf("%d ", v)
        if (i+1)%10 == 0 {
            fmt.Println()
        }
    }
    fmt.Println()
}