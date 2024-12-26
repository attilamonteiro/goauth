package services

import (
    "encoding/json"
    "log"
    "net/http"

    "golang.org/x/crypto/bcrypt"

    "goauth/models"
    "goauth/repository"
    "goauth/utils"
)

// Login faz a autenticação do usuário, gera o token e o retorna em JSON.
func Login(requestUser *models.User) (int, []byte) {
    token, err := authenticateUser(requestUser.Username, requestUser.Password)
    if err != nil {
        log.Println("Authentication error:", err)
        return http.StatusUnauthorized, []byte("Invalid credentials")
    }

    response, _ := json.Marshal(map[string]string{"token": token})
    return http.StatusOK, response
}

// Register cria um novo usuário no banco (senha com hash bcrypt).
func Register(newUser *models.User) (int, []byte) {
    err := registerUser(newUser.Username, newUser.Password)
    if err != nil {
        log.Println("Registration error:", err)
        return http.StatusInternalServerError, []byte("Error creating user in DB")
    }

    response, _ := json.Marshal(map[string]string{"message": "User registered successfully"})
    return http.StatusOK, response
}

// RefreshToken gera um novo token com base no UUID do usuário informado.
func RefreshToken(requestUser *models.User) (int, []byte) {
    token, err := utils.GenerateToken(requestUser.UUID)
    if err != nil {
        return http.StatusInternalServerError, []byte("Error generating token")
    }

    response, _ := json.Marshal(map[string]string{"token": token})
    return http.StatusOK, response
}

// Logout é um stub para invalidar o token, se desejar (ex.: usar blacklist, Redis etc.).
func Logout(req *http.Request) error {
    return nil
}

func authenticateUser(username, password string) (string, error) {
    user, err := repository.FindByUsername(username)
    if err != nil {
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", err
    }

    token, err := utils.GenerateToken(user.UUID)
    if err != nil {
        return "", err
    }

    return token, nil
}

func registerUser(username, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user := &models.User{
        Username: username,
        Password: string(hashedPassword),
    }

    return repository.CreateUser(user)
}
