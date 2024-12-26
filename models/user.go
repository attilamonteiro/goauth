package models

// Definindo o modelo User
type User struct {
    UUID     string `json:"uuid"`
    Username string `json:"username"`
    Password string `json:"password"`
}
