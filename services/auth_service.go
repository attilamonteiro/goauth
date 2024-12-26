package services

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"goauth/models"
	"goauth/repository"
	"goauth/core"
)

// Login autentica o usuário e retorna um token JWT.
func Login(userRepo repository.UserRepository, requestUser *models.User) (int, []byte) {
	// Verifica credenciais e gera token
	token, err := authenticateUser(userRepo, requestUser.Username, requestUser.Password)
	if err != nil {
		log.Printf("Authentication error for username %s: %v", requestUser.Username, err)
		return http.StatusUnauthorized, []byte(`{"error": "Invalid credentials"}`)
	}

	// Retorna o token no formato JSON
	response, _ := json.Marshal(map[string]string{"token": token})
	return http.StatusOK, response
}

// Register cria um novo usuário com senha hash.
func Register(userRepo repository.UserRepository, newUser *models.User) (int, []byte) {
	// Verifica se o nome de usuário já existe
	if existingUser, _ := userRepo.FindByUsername(newUser.Username); existingUser != nil {
		log.Printf("Attempted registration with existing username: %s", newUser.Username)
		return http.StatusConflict, []byte(`{"error": "Username already exists"}`)
	}

	// Cria o novo usuário
	if err := registerUser(userRepo, newUser.Username, newUser.Password); err != nil {
		log.Printf("Registration error for username %s: %v", newUser.Username, err)
		return http.StatusInternalServerError, []byte(`{"error": "Error creating user in DB"}`)
	}

	// Retorna mensagem de sucesso
	response, _ := json.Marshal(map[string]string{"message": "User registered successfully"})
	return http.StatusCreated, response
}

// RefreshToken gera um novo token JWT para o usuário.
func RefreshToken(requestUser *models.User) (int, []byte) {
	// Gera um novo token JWT
	token, err := auth.GenerateToken(requestUser.UUID)
	if err != nil {
		log.Printf("Token generation error for user UUID %s: %v", requestUser.UUID, err)
		return http.StatusInternalServerError, []byte(`{"error": "Error generating token"}`)
	}

	// Retorna o token no formato JSON
	response, _ := json.Marshal(map[string]string{"token": token})
	return http.StatusOK, response
}

// Logout é um stub para invalidar o token, se necessário.
func Logout(req *http.Request) error {
	// Implementação futura, ex.: invalidar token em uma blacklist
	return nil
}

// authenticateUser verifica as credenciais do usuário e gera um token JWT.
func authenticateUser(userRepo repository.UserRepository, username, password string) (string, error) {
	// Busca o usuário pelo nome
	user, err := userRepo.FindByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) || user == nil {
			return "", errors.New("user not found")
		}
		return "", err
	}

	// Compara a senha fornecida com o hash armazenado
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Gera um token JWT
	return auth.GenerateToken(user.UUID)
}

// registerUser cria um novo usuário com a senha hash.
func registerUser(userRepo repository.UserRepository, username, password string) error {
	// Gera o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Cria a estrutura do usuário
	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	// Salva o usuário no banco de dados
	return userRepo.CreateUser(user)
}
