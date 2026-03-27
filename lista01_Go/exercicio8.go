package main 

import "fmt"

const Pi float64 = 3.141592
var R, H, A, C float64

func main() {
	fmt.Println("Qual a medida do raio da latinha, em metros?")
	fmt.Scan(&R)
	fmt.Println("Qual a medida da altura da latinha, em metros?")
	fmt.Scan(&H)
	A = (2*Pi*R*(R+H))
	fmt.Println("A área total do cilindro é", A)
	C = (A*100)
	fmt.Println("O custo total é", C)
}