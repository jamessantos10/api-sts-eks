package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

// HealthStatus Estrutura para representar o status de saúde do serviço
type HealthStatus struct {
	Status string `json:"status"`
}

// healthCheckHandler Rota para verificar o status de saúde do serviço
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o serviço está em execução corretamente
	// Você pode adicionar lógica adicional aqui, se necessário

	// Cria uma estrutura de status de saúde
	status := HealthStatus{
		Status: "UP",
	}

	// Serializa a estrutura de status para JSON
	response, err := json.Marshal(status)
	if err != nil {
		log.Println("Error > healthCheckHandler >", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho de resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Escreve a resposta
	w.Write(response)
}
