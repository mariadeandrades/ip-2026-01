package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"github.com/servidorHTTP/servidorHTTP/app/handlers"
	"github.com/servidorHTTP/servidorHTTP/app/utils"
)

func main() {
	utils.ConnectToDB()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./app/static"))))

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/createPacientes", handlers.CreatePatientHandler)
	http.HandleFunc("/pacientes", handlers.ListPatientsHandler)
	http.HandleFunc("/updatePacientes", handlers.UpdatePatientHandler)
	http.HandleFunc("/deletePacientes", handlers.DeletePatientHandler)

	localIP := getLocalIP()

	fmt.Println("Servidor rodando em:")
	fmt.Println("Local: http://localhost:8080")
	fmt.Printf("Rede: http://%s:8080\n", localIP)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}

	return "127.0.0.1"
}
