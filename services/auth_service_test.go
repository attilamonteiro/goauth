package services

import (
	"goauth/config"
	"goauth/models"
	"goauth/repository"
	"testing"
)

func TestLogin(t *testing.T) {
	// Inicializa a conexão com o banco de dados
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// Inicializa o repositório
	userRepo := repository.NewUserRepository(db)

	// Cria um usuário de teste
	testUser := &models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	err := userRepo.CreateUser(testUser)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	// Testa login com credenciais corretas
	status, _ := Login(userRepo, testUser)
	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

	// Testa login com credenciais incorretas
	testUser.Password = "wrongpassword"
	status, _ = Login(userRepo, testUser)
	if status != 401 {
		t.Errorf("Expected status 401, got %d", status)
	}
}

func TestRegister(t *testing.T) {
	// Inicializa a conexão com o banco de dados
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// Inicializa o repositório
	userRepo := repository.NewUserRepository(db)

	// Cria um usuário de teste
	testUser := &models.User{
		Username: "newuser",
		Password: "newpassword",
	}

	// Testa registro de usuário
	status, _ := Register(userRepo, testUser)
	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}

	// Testa registro com nome de usuário existente
	status, _ = Register(userRepo, testUser)
	if status != 500 {
		t.Errorf("Expected status 500, got %d", status)
	}
}

func TestRefreshToken(t *testing.T) {
	// Inicializa a conexão com o banco de dados
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	// Inicializa o repositório
	userRepo := repository.NewUserRepository(db)

	// Cria um usuário de teste
	testUser := &models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	err := userRepo.CreateUser(testUser)
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	// Testa refresh token
	status, _ := RefreshToken(testUser)
	if status != 200 {
		t.Errorf("Expected status 200, got %d", status)
	}
}

func TestLogout(t *testing.T) {
	// Testa logout (atualmente um stub)
	err := Logout(nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
