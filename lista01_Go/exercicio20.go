/*algoritmo_segundos

declare horas, minutos, segundos, horaseg, minutoseg, segundostotal numerico

escreva ("Informe o horário na seguinte ordem: horas, em horas; minutos, em minutos; segundos, em segundos.")
receba horas minutos segundos
funcao horaseg(horas)
    escreva (horas*3600)
funcao minutoseg(minutos)
    escreva (minutos*60)
segundostotal = (horaseg + minutoseg + segundos)
escreval ("O tempo em segundo é", segundostotal ".")

fim_algoritmo_segundos*/
package main

import "fmt"

func main() {
	var horas, minutos, segundos, horaseg, minutoseg, segundostotal float32
	fmt.Println("Informe o tempo transcorrido na seguinte ordem: as horas, em horas; os minutos, em minutos; os segundos, em segundos.")
	fmt.Scan(&horas, &minutos, &segundos)
	horaseg = horas*3600
	minutoseg = minutos*60
	segundostotal = horaseg + minutoseg + segundos
	fmt.Printf("A quantidade de segundos é %F.", segundostotal)
}