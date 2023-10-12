package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Json retorna uma reposta em json para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)
	if dados != nil {

		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}

	}
}

// Erro retorna um erro em json para a requisição
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
