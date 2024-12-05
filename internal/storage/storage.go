package storage

import (
	"errors"
	"sync"

	"github.com/JeanGrijp/meu-primeiro-crud-go/internal/models"
	"github.com/google/uuid"
)

type Storage struct {
	data map[uuid.UUID]models.User // Mapa em memória que funciona como banco de dados
	mu   sync.Mutex                // Mutex para garantir segurança em operações concorrentes
}

// Função para inicializar o "banco de dados"
func NewStorage() *Storage {
	return &Storage{
		data: make(map[uuid.UUID]models.User),
	}
}

func (s *Storage) FindAll() []models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	users := make([]models.User, 0, len(s.data)) // Cria um slice vazio
	for _, user := range s.data {                // Itera sobre o mapa
		users = append(users, user)
	}
	return users
}

func (s *Storage) FindById(id uuid.UUID) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.data[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *Storage) Insert(user models.User) models.User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = uuid.New()
	s.data[user.ID] = user
	return user
}

func (s *Storage) Update(id uuid.UUID, updated models.User) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.data[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}

	updated.ID = id      // Mantém o mesmo ID
	s.data[id] = updated // Atualiza os dados no mapa
	return updated, nil
}

func (s *Storage) Delete(id uuid.UUID) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.data[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}

	delete(s.data, id) // Remove o usuário do mapa
	return user, nil
}
