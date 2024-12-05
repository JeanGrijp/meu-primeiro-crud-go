package handlers

import (
	"net/http"

	"github.com/JeanGrijp/meu-primeiro-crud-go/internal/models"
	"github.com/JeanGrijp/meu-primeiro-crud-go/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUserHandler(store *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		// Lê o JSON do corpo da requisição
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid user data",
				"error":   err.Error(),
			})
			return
		}

		// Valida os dados do usuário
		if validationErr := user.Validate(); validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid user data",
				"error":   validationErr.Error(),
			})
			return
		}

		// Insere o novo usuário no "banco de dados"
		newUser := store.Insert(user)
		// Retorna o usuário criado
		c.JSON(http.StatusCreated, newUser)
	}
}

func GetUsersHandler(store *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Busca todos os usuários do "banco de dados"
		users := store.FindAll()
		// Retorna a lista de usuários (ou vazia)
		c.JSON(http.StatusOK, users)
	}
}

func GetUserHandler(store *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o ID do parâmetro da URL
		idParam := c.Param("id")
		id, err := uuid.Parse(idParam) // Tenta converter o ID para UUID
		if err != nil {
			// Retorna erro se o ID for inválido
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid user ID",
				"error":   err.Error(),
			})
			return
		}

		// Busca o usuário no "banco de dados"
		user, err := store.FindById(id)
		if err != nil {
			// Retorna erro se o usuário não for encontrado
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}

		// Retorna o único usuário encontrado
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUserHandler(store *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o ID da URL
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
			return
		}

		// Lê o JSON do corpo da requisição
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil || user.Validate() != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user data"})
			return
		}

		// Atualiza o usuário no "banco de dados"
		updatedUser, err := store.Update(id, user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}

		// Retorna o usuário atualizado
		c.JSON(http.StatusOK, updatedUser)
	}
}

func DeleteUserHandler(store *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o ID da URL
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user ID"})
			return
		}

		// Remove o usuário do "banco de dados"
		user, err := store.Delete(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return
		}

		// Retorna o usuário deletado
		c.JSON(http.StatusOK, user)
	}
}
