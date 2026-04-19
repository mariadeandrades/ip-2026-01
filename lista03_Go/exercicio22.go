package main

import "fmt"

func main() {
    var S float64 = 0

    j := 37.0
    for i := 1.0; i <= 37.0; i ++{ //i é o denominador, j é o numerador que está mais à esquerda, q é sempre 1 unidade menor q o numerador à direita
        termo := (j * (j + 1)) / i  //termo atual
        S += termo
        j -- //ex: 35*36 + 34*35 + 33*34 ...
    }
    fmt.Printf("S = %.2f", S)
}