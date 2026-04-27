package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Funcionou! Voce acessou: %s", r.URL.Path)
}

func main() {
	fmt.Println("Tentando iniciar o servidor na porta 8080...")
	
	http.HandleFunc("/", handler)
	
	// Guardamos o erro em uma variável para analisar
	err := http.ListenAndServe(":8081", nil)
	
	if err != nil {
		fmt.Printf("ERRO FATAL: %v\n", err)
		log.Fatal(err)
	}
}