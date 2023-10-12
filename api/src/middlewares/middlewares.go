package middlewares

import (
	autenticacao "api/src/authentication"
	respostas "api/src/responses"
	"log"
	"net/http"
)

func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Autenticar verifica se o usuario está fazendo a requisição autenticando
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		proximaFuncao(w, r)
	}
}
