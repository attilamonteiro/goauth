package controllers

import (
    "encoding/json"
    "goauth/services"
    "goauth/models"
    "net/http"
)

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    var requestUser models.User
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&requestUser); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Verifica se é um login mockado
    if requestUser.Username == "mockuser" && requestUser.Password == "mockpassword" {
        mockToken := map[string]string{"token": "mocked_jwt_token"}
        responseData, _ := json.Marshal(mockToken)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(responseData)
        return
    }

    // Chama o serviço de login
    status, token := services.Login(&requestUser)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(token)
}

// Register godoc
// @Summary Register user
// @Description Create a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "New user data"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    var newUser models.User
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&newUser); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Chama o serviço de registro
    status, response := services.Register(&newUser)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    var requestUser models.User
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&requestUser); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Chama o serviço para gerar um novo token
    status, token := services.RefreshToken(&requestUser)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(token)
}

func Logout(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // Chama o serviço de logout
    err := services.Logout(r)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
    } else {
        w.WriteHeader(http.StatusOK)
    }
}
