package main

import (
	"api-sts-eks/api/src/router"
	"fmt"
	"log"
	"net/http"
)

// Rodando API na porta 8080
func main() {
	fmt.Println("Rodando API")
	r := router.Gerar()
	log.Fatal(http.ListenAndServe(":8080", r))
}
