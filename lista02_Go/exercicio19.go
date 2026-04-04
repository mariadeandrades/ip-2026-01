package main 

import (
	"fmt"
	"math"
)

var novo float64

func raiz(rc, hc float64) float64 {
	return math.Sqrt((rc*rc + hc*hc))
}

func main() {
	var a, v, hc, rc, hi, ri, re float64
	var n int
	const p float64 = 3.14
	fmt.Println("Deseja calcular a área superficial e o volume de qual sólido geométrico: cone, cilindro ou esfera?")
	fmt.Println("Digite o número correspondente ao sólido escolhido, de acordo com a seguinte lista:")
	fmt.Println("(1) : Cone\n(2) : Cilindro\n(3) : Esfera")
	fmt.Scan(&n)
	switch n{
	case 1:
		fmt.Println("Insira os valores do raio da base e da altura do cone, respectivamente, em metros.")
		fmt.Scan(&rc, &hc)
		novo := raiz(rc, hc)
		a = (p*rc*novo)
		v = ((p*rc*rc*hc)/3)
		fmt.Printf("A área superficial do cone é %.2f e seu volume é %.2f.\n", a, v)
	case 2:
		fmt.Println("Insira os valores do raio da base e da altura do cilindro, respectivamente, em metros.")
		fmt.Scan(&ri, &hi)
		a = (2*p*ri*(ri + hi))
		v = (p*ri*ri*hi)
		fmt.Printf("A área superficial do cilindro é %.2f e seu volume é %.2f.\n", a, v)
	case 3:
		fmt.Println("Insira o valor do raio da esfera, em metros.")
		fmt.Scan(&re)
		a = (4*p*re*re)
		v = ((4*p*re*re*re)/3)
		fmt.Printf("A área superficial da esfera é %.2f e seu volume é %.2f.\n", a, v)
	default:
		fmt.Println("Opção inválida.")
	}
}

