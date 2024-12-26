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
    // 1. Busca o usuário no banco pelo username
    userFromDb, err := repository.FindByUsername(requestUser.Username)
    if err != nil || userFromDb == nil {
        log.Println("User not found or error:", err)
        return http.StatusUnauthorized, []byte("Invalid credentials")
    }

    // 2. Compara a senha informada com o hash salvo
    if err := bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(requestUser.Password)); err != nil {
        log.Println("Password mismatch:", err)
        return http.StatusUnauthorized, []byte("Invalid credentials")
    }

    // 3. Gera um token JWT (HS256) usando a função em utils
    token, err := utils.GenerateToken(userFromDb.UUID)
    if err != nil {
        log.Println("Error generating token:", err)
        return http.StatusInternalServerError, []byte("Error generating token")
    }

    // 4. Retorna o token em JSON
    response, _ := json.Marshal(map[string]string{"token": token})
    return http.StatusOK, response
}

// Register cria um novo usuário no banco (senha com hash bcrypt).
func Register(newUser *models.User) (int, []byte) {
    // 1. Gera o hash da senha
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err != nil {
        return http.StatusInternalServerError, []byte("Error hashing password")
    }
    newUser.Password = string(hashedPassword)

    // 2. Salva o novo usuário no banco
    // (Certifique-se de ter a função CreateUser no seu repository)
    err = repository.CreateUser(newUser)
    if err != nil {
        return http.StatusInternalServerError, []byte("Error creating user in DB")
    }

    // 3. Retorna resposta de sucesso
    response, _ := json.Marshal(map[string]string{"message": "User registered successfully"})
    return http.StatusOK, response
}

// RefreshToken gera um novo token com base no UUID do usuário informado.
func RefreshToken(requestUser *models.User) (int, []byte) {
    // Exemplo de revalidação: pode checar se o usuário existe no DB, etc.
    // Abaixo, apenas gera outro token diretamente:
    token, err := utils.GenerateToken(requestUser.UUID)
    if err != nil {
        return http.StatusInternalServerError, []byte("Error generating token")
    }

    response, _ := json.Marshal(map[string]string{"token": token})
    return http.StatusOK, response
}

// Logout é um stub para invalidar o token, se desejar (ex.: usar blacklist, Redis etc.).
func Logout(req *http.Request) error {
    // Exemplo simplificado: não faz nada
    return nil
}

func Authenticate(username, password string) (string, error) {
    user, err := repository.FindByUsername(username)
    if err != nil {
        return "", err
    }

    // Compare the provided password with the stored hash
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", err
    }

    // Generate a token
    token, err := utils.GenerateToken(user.UUID)
    if err != nil {
        return "", err
    }

    return token, nil
}

func RegisterUser(username, password string) error {
    // Hash the password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Create the user object
    user := &models.User{
        Username: username,
        Password: string(hashedPassword),
    }

    // Save the user to the repository
    err = repository.CreateUser(user)
    if err != nil {
        return err
    }

    return nil
}
