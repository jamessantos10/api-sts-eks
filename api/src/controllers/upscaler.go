package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// caatalog Estrutura para receber os valores da rota /up
type catalog struct {
	Power     bool   `json:"power"`
	Idaccount string `json:"idaccount"`
}

// Receber parametro para ligar o cluster
func Upscaler(w http.ResponseWriter, r *http.Request) {

	// Lê o corpo do Post recebido, usando o pacote ioutil e armazena na variável corpoRequisicao
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	// Se houver erro durante a leitura do corpoRequisicao, retorne:
	if erro != nil {
		w.Write([]byte("Falha ao ler corpo da requisição"))
		return
	}

	// Contém os dados em Json
	var catalog catalog

	// Converte os dados json corpoRequisicao para a estrutura do catalog
	if erro = json.Unmarshal(corpoRequisicao, &catalog); erro != nil {
		w.Write([]byte("Erro ao receber parametro"))
		return
	}

	fmt.Println(catalog)

	// Se o parametro recebido for true, chame a função assumerolests
	if catalog.Power == true {
		{
			assumerolests(catalog.Idaccount)
		}
	}
}
