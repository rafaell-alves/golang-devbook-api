package controllers

import (
	autenticacao "api/src/authentication"
	banco "api/src/database"
	modelos "api/src/models"
	repositorios "api/src/repository"
	respostas "api/src/responses"
	seguranca "api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuario.Senha, usuarioSalvoNoBanco.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	token, _ := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)

	w.Write([]byte(token))
}
