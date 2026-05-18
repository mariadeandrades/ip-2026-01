package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/lib/pq"
)

func ConnectToDB() {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "Mel2026#")
	dbname := getEnv("DB_NAME", "first_teste")

	ensureDatabase(host, port, user, password, dbname)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname,
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com banco:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	ensurePacientesTable()

	fmt.Println("Conectado ao banco de dados com sucesso!")
}

func ensureDatabase(host string, port string, user string, password string, dbname string) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		host,
		port,
		user,
		password,
	)

	adminDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao abrir conexao administrativa:", err)
	}
	defer adminDB.Close()

	if err := adminDB.Ping(); err != nil {
		log.Fatal("Erro ao conectar no banco postgres:", err)
	}

	_, err = adminDB.Exec(`CREATE DATABASE ` + pq.QuoteIdentifier(dbname))
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "42P04" {
			return
		}

		log.Fatal("Erro ao criar banco de dados:", err)
	}
}

func ensurePacientesTable() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS pacientes (
			id SERIAL PRIMARY KEY,
			nome VARCHAR(100) NOT NULL,
			cpf VARCHAR(20) NOT NULL,
			idade INT NOT NULL,
			peso DECIMAL(4,1) NOT NULL,
			tipo_sangue VARCHAR(3) NOT NULL
		)
	`)
	if err != nil {
		log.Fatal("Erro ao criar tabela pacientes:", err)
	}
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
