package main 

import (
   "fmt"
   "strings"
)

func main() {
   var destino, tipo string

   fmt.Println("Qual a região de destino da sua viagem (norte, nordeste, sul ou centro-oeste)?")
   fmt.Scan(&destino)
   fmt.Println("Qual tipo de viagem você realizará (ida ou completa, ou seja, ida e volta)?")
   fmt.Scan(&tipo)

   destino = strings.ToLower(strings.TrimSpace(destino))
   tipo = strings.ToLower(strings.TrimSpace(tipo))

   switch destino {
   case "norte":
      if tipo == "ida"{
         fmt.Println("Sua viagem custará R$ 500,00.")
      }else if tipo == "completa"{
         fmt.Println("Sua viagem custará R$ 900,00.")
      }else{
         fmt.Println("Tipo de viagem inválido.")
      }
   case "nordeste":
      if tipo == "ida"{
          fmt.Println("Sua viagem custará R$ 350,00.")
      }else if tipo == "completa"{
         fmt.Println("Sua viagem custará R$ 650,00.")
      }else{
         fmt.Println("Tipo de viagem inválido.")
      }
   case "centro-oeste":
      if tipo == "ida"{
       fmt.Println("Sua viagem custará R$ 350,00.")
      }else if tipo == "completa"{
         fmt.Println("Sua viagem custará R$ 600,00.")
      }else{
         fmt.Println("Tipo de viagem inválido.")
      }
   case "sul":
      if tipo == "ida"{
          fmt.Println("Sua viagem custará R$ 300,00.")
      }else if tipo == "completa"{
         fmt.Println("Sua viagem custará R$ 550,00.")
      }else{
         fmt.Println("Tipo de viagem inválido.")
      }
   default:
      fmt.Println("Digite um destino válido.")
   } 
}