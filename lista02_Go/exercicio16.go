package main 

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var a, b, c float64
	var x1, x2 complex128
	var delta complex128
	fmt.Println("Este programa utiliza a seguinte equação quadrática: ( A*x*x + B*x + C = 0).\nDigite os valores dos coeficientes A, B e C.")
	fmt.Scan(&a, &b, &c)
	if a == 0 {
		fmt.Println("Coeficiente A não pode ser zero para uma equação quadrática.")
		return
	}
	delta = complex(b*b-4*a*c, 0)
	w := cmplx.Sqrt(delta)

	x1 = (-complex(b, 0) + w) / complex(2*a, 0)
	x2 = (-complex(b, 0) - w) / complex(2*a, 0)

	if real(delta) > 0 && imag(delta) == 0 {
		fmt.Printf("As raizes são %.2f e %.2f.\n", real(x1), real(x2))
	} else if real(delta) == 0 && imag(delta) == 0 {
		fmt.Printf("As duas raízes são iguais, e seu valor é %.2f.\n", real(x1))
	} else {
		fmt.Printf("Há duas raízes complexas, e seus valores são %v e %v.\n", x1, x2)
	}
}

