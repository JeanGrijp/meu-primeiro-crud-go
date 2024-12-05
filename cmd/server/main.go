package main

import (
	"github.com/JeanGrijp/meu-primeiro-crud-go/internal/handlers"
	"github.com/JeanGrijp/meu-primeiro-crud-go/internal/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	store := storage.NewStorage() // Inicializa o banco de dados
	r := gin.Default()

	// Rotas
	api := r.Group("/api/users") // Prefixo comum para todas as rotas de usuários
	{
		api.POST("", handlers.CreateUserHandler(store))           // Criação de usuário
		api.GET("", handlers.GetUsersHandler(store))              // Listar todos os usuários ou buscar por ID (query string)
		api.GET("unic/:id", handlers.GetUserHandler(store))       // Buscar usuário por ID (parâmetro na URL)
		api.PUT("unic/:id", handlers.UpdateUserHandler(store))    // Atualizar usuário por ID
		api.DELETE("unic/:id", handlers.DeleteUserHandler(store)) // Deletar usuário por ID
	}

	r.Run(":8080")
}
