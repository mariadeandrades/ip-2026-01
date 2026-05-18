package handlers

import (
	"html/template"
	"net/http"
	"github.com/servidorHTTP/servidorHTTP/app/utils"
	"strconv"
)

type Patient struct {
	ID          int
	Nome        string
	CPF         string
	Idade       int
	Peso    	float64
	TipoSangue  string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/static/index.html")
	if err != nil {
		http.Error(w, "Erro ao carregar página inicial", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("app/static/forms/createPatient.html")
		if err != nil {
			http.Error(w, "Erro ao carregar formulário", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		nome := r.FormValue("nome")
		cpf := r.FormValue("cpf")
		idade, _ := strconv.Atoi(r.FormValue("idade"))
		peso, _ := strconv.ParseFloat(r.FormValue("peso"), 64)
		tipoSangue := r.FormValue("tipo_sangue")

		_, err := utils.DB.Exec(`
			INSERT INTO patients (nome, cpf, idade, peso, tipo_sangue)
			VALUES ($1, $2, $3, $4, $5)
		`, nome, cpf, idade, peso, tipoSangue)

		if err != nil {
			http.Error(w, "Erro ao cadastrar paciente: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/pacientes", http.StatusSeeOther)
	}
}

func ListPatientsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := utils.DB.Query(`
		SELECT id, nome, cpf, idade, peso, tipo_sangue
		FROM pacientes
		ORDER BY id
	`)

	if err != nil {
		http.Error(w, "Erro ao buscar pacientes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var patients []Patient

	for rows.Next() {
		var p Patient

		err := rows.Scan(
			&p.ID,
			&p.Nome,
			&p.CPF,
			&p.Idade,
			&p.Peso,
			&p.TipoSangue,
		)

		if err != nil {
			http.Error(w, "Erro ao ler paciente: "+err.Error(), http.StatusInternalServerError)
			return
		}

		patients = append(patients, p)
	}

	tmpl, err := template.ParseFiles("app/static/forms/listPatients.html")
	if err != nil {
		http.Error(w, "Erro ao carregar lista", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, patients)
}

func UpdatePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")

		var p Patient

		err := utils.DB.QueryRow(`
			SELECT id, nome, cpf, idade, peso, tipo_sangue
			FROM pacientes
			WHERE id = $1
		`, id).Scan(
			&p.ID,
			&p.Nome,
			&p.CPF,
			&p.Idade,
			&p.Peso,
			&p.TipoSangue,
		)

		if err != nil {
			http.Error(w, "Paciente não encontrado: "+err.Error(), http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles("app/static/forms/updatePatient.html")
		if err != nil {
			http.Error(w, "Erro ao carregar formulário de atualização", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, p)
		return
	}

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		cpf := r.FormValue("cpf")
		idade, _ := strconv.Atoi(r.FormValue("idade"))
		peso, _ := strconv.ParseFloat(r.FormValue("peso"), 64)
		tipoSangue := r.FormValue("tipo_sangue")

		_, err := utils.DB.Exec(`
			UPDATE pacientes
			SET nome = $1,
				cpf = $2,
				idade = $3,
				peso = $4,
				tipo_sangue = $5
			WHERE id = $6
		`, nome, cpf, idade, peso, tipoSangue, id)

		if err != nil {
			http.Error(w, "Erro ao atualizar paciente: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/pacientes", http.StatusSeeOther)
	}
}

func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "ID do paciente não informado", http.StatusBadRequest)
		return
	}

	_, err := utils.DB.Exec("DELETE FROM pacientes WHERE id = $1", id)

	if err != nil {
		http.Error(w, "Erro ao excluir paciente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pacientes", http.StatusSeeOther)
}
