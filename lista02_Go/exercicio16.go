/* Desenvolver um programa que leia os coeficientes (A, B e C) de uma equação do segundo grau ( A*x*x
+ B*x + C
=0) e que calcule suas raízes. O programa deve mostrar, quando possível, o valor das raízes calculadas e a
classificação das mesmas: “RAÍZES IMAGINÁRIAS”, “RAIZ ÚNICA” ou “RAÍZES DISTINTAS”.
*/


//vou terminar depois, não consigo raciocinar mais 

package main 

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	var x1, x2  complex64
	var delta, w float64
	fmt.Println("Este programa utiliza a seguinte equação quadrática: ( A*x*x + B*x + C = 0).\nDigite os valores dos coeficientes A, B e C.")
	fmt.Scan(&a, &b, &c)
	delta = (b*b - 4*a*c)
	w = math.Sqrt(delta)

	if delta > 0{
		x1 = ((-b + w)/ 2*a)
		x2 = ((-b - w)/ 2*a)
		fmt.Printf("As raizes são %v e %v.", x1, x2)
	}else if delta == 0{
		x1 = (-b / 2*a)
		fmt.Printf("As duas raízes são iguais, e seu valor é %v.", x1)
	}else{
		x1 = ((-b + w)/ 2*a)
		x2 = ((-b - w)/ 2*a)
		fmt.Println("Há duas raízes complexas, e seus valores são %v e %v.", x1, x2)
	}
}

