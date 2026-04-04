/* Desenvolver um programa que leia os coeficientes (A, B e C) de uma equação do segundo grau ( A*x*x
+ B*x + C
=0) e que calcule suas raízes. O programa deve mostrar, quando possível, o valor das raízes calculadas e a
classificação das mesmas: “RAÍZES IMAGINÁRIAS”, “RAIZ ÚNICA” ou “RAÍZES DISTINTAS”.
*/

package main 

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	var x complex64
	fmt.Println("Este programa utiliza a seguinte equação quadrática: ( A*x*x + B*x + C = 0).\nDigite os valores dos coeficientes A, B e C.")
	fmt.Scan(&a, &b, &c)
	x = calculo(a, b, c)
	if x 
	fmt.Printf("Os resultados da equação são: %v.")
}

func calculo(a, b, c int) complex64 {
	return ((a*x*x + b*x + c) = 0)
}