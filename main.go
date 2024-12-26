package main

import (
    "fmt"
    "log"
    "net/http"
    "goauth/routers"
    "goauth/repository"

    "github.com/gorilla/mux"
    httpSwagger "github.com/swaggo/http-swagger" // Adicionado
    _ "goauth/docs" // Adicionado
)

// @title GoAuth API
// @version 1.0
// @description This is a sample server for GoAuth.
// @host localhost:8080
// @BasePath /
func main() {
    // Initialize the database connection
    repository.InitDB("goauth.db")

    // Inicializando o roteador
    router := mux.NewRouter()
    
    // Define as rotas de autenticação (login, register etc.)
    routers.SetAuthenticationRoutes(router)

    // Adicionando a rota do Swagger
    // Isso faz com que ao acessar /swagger/*, seja servido o swagger UI
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    // Inicializando o servidor
    fmt.Println("API rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
