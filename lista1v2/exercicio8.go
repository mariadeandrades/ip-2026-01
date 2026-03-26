/*
algoritmo_custo_aluminio

const real P = 3.14159
declare R, H, V float

escreva ("Qual a medida do raio da latinha, em metros?")
receba R
escreva ("Qual a medida da altura da latinha, em metros?")
receba H

#função A que calcula a área total do cilindro
funcao A(P, R, H)
    devolva (2*P*R(R+H))

#função que calcula o custo relacionado à área da lata e ao valor do alumínio
funcao C(A)
    devolva (A*100)

escreval ("O custo para a produção da latinha é de", C " reais.")

fim_algoritmo_custo_aluminio
*/

package main 

import "fmt"

const Pi float64 = 3.14159
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