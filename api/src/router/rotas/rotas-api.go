package rotas

import (
	"net/http"
	"api-sts-eks/api/src/controllers"
)

// Rotas expostas
var rotasApi = []Rota{
	// Receber os parametros do Catalogo
	{
		URI:                "/up",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Upscaler,
		RequerAutenticacao: false,
	},
	// Validar o Health Check
	{
		URI:                "/health",
		Metodo:             http.MethodGet,
		Funcao:             controllers.HealthCheckHandler,
		RequerAutenticacao: false,
	},
}
