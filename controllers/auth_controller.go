package controllers

import (
	"encoding/json"
	"goauth/models"
	"goauth/repository"
	"goauth/services"
	"net/http"
)

// respondJSON envia a resposta JSON padronizada
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		json.NewEncoder(w).Encode(payload)
	}
}

// respondError envia uma resposta de erro padronizada
func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

// Login autentica um usuário e retorna um token JWT.
func Login(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var requestUser models.User
	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Chamar o serviço de autenticação
	status, response := services.Login(userRepo, &requestUser)

	// Responder ao cliente
	if status == http.StatusUnauthorized {
		respondError(w, status, "Invalid username or password")
		return
	}
	respondJSON(w, status, response)
}

// Register registra um novo usuário no sistema.
func Register(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Chamar o serviço de registro
	status, response := services.Register(userRepo, &newUser)

	// Responder ao cliente
	if status == http.StatusConflict {
		respondError(w, status, "Username already exists")
		return
	}
	respondJSON(w, status, response)
}

// RefreshToken gera um novo token JWT para o usuário autenticado.
func RefreshToken(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var requestUser models.User
	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Chamar o serviço de refresh token
	status, response := services.RefreshToken(&requestUser)
	respondJSON(w, status, response)
}

// Logout realiza o logout do usuário.
func Logout(w http.ResponseWriter, r *http.Request, userRepo repository.UserRepository) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Chamar o serviço de logout
	if err := services.Logout(r); err != nil {
		respondError(w, http.StatusInternalServerError, "Failed to logout")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Logged out successfully"})
}
