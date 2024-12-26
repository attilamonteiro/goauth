package main

import (
	"fmt"
	"goauth/config"
	"goauth/repository"
	"goauth/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "goauth/docs" 
)

func main() {
	// Configurar conexão com o banco de dados
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// Inicializar o repositório
	userRepo := repository.NewUserRepository(db)

	// Configurar rotas
	router := mux.NewRouter()

	// Rotas de autenticação
	routers.SetAuthenticationRoutes(router, userRepo)

	// Rota para o Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Iniciar servidor
	fmt.Println("API rodando em http://localhost:8080")
	fmt.Println("Swagger da API rodando em http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", router))
}
