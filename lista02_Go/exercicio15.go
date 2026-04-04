// receber três valores inteiros (dia, mês e ano) e converta mês pra outra variável do tipo string, depois concatene tudo
package main 

import (
	"fmt"
)

var dia, mes, ano int

func main() {
	fmt.Println("Digite a data numericamente no seguinte formato e ordem:\ndia mês ano")
	fmt.Scan(&dia, &mes, &ano)

switch mes {
case 1:
	fmt.Printf("A data por extenso é %v de Janeiro de %v.", dia, ano)
case 2:
	fmt.Printf("A data por extenso é %v de Fevereiro de %v.", dia, ano)
case 3:
	fmt.Printf("A data por extenso é %v de Março de %v.", dia, ano)
case 4:
	fmt.Printf("A data por extenso é %v de Abril de %v.", dia, ano)
case 5:
	fmt.Printf("A data por extenso é %v de Maio de %v.", dia, ano)
case 6:
	fmt.Printf("A data por extenso é %v de Junhoo de %v.", dia, ano)
case 7:
	fmt.Printf("A data por extenso é %v de Julho de %v.", dia, ano)
case 8:
	fmt.Printf("A data por extenso é %v de Agosto de %v.", dia, ano)
case 9:
	fmt.Printf("A data por extenso é %v de Setembro de %v.", dia, ano)
case 10:
	fmt.Printf("A data por extenso é %v de Outubro de %v.", dia, ano)
case 11:
	fmt.Printf("A data por extenso é %v de Novembro de %v.", dia, ano)
case 12:
	fmt.Printf("A data por extenso é %v de Dezembro de %v.", dia, ano)
default:
	fmt.Println("Data inválida.")
}

}