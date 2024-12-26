package models

// Definindo o modelo User
type User struct {
	UUID     string `gorm:"primaryKey" json:"uuid"`      // Define UUID como chave primária
	Username string `gorm:"unique;not null" json:"username"` // Username deve ser único e não pode ser nulo
	Password string `gorm:"not null" json:"password"`    // Password não pode ser nulo
}
