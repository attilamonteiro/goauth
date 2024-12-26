package repository

import (
	"goauth/config"
	"goauth/models"
	"testing"
)

func TestUserRepository(t *testing.T) {
	// Inicializa a conexão com o banco de dados
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// Inicializa o repositório
	userRepo := NewUserRepository(db)

	// Cria um usuário de teste
	testUser := &models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	err := userRepo.CreateUser(testUser)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Testa busca por nome de usuário
	user, err := userRepo.FindByUsername(testUser.Username)
	if err != nil {
		t.Fatalf("Failed to find user: %v", err)
	}
	if user.Username != testUser.Username {
		t.Errorf("Expected username %s, got %s", testUser.Username, user.Username)
	}
}

func TestUserRepositoryDuplicate(t *testing.T) {
	// Inicializa a conexão com o banco de dados
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// Inicializa o repositório
	userRepo := NewUserRepository(db)

	// Cria um usuário de teste
	testUser := &models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	err := userRepo.CreateUser(testUser)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Tenta criar o mesmo usuário novamente
	err = userRepo.CreateUser(testUser)
	if err == nil {
		t.Fatalf("Expected error when creating duplicate user, got nil")
	}
}
