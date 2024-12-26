package repository

import (
	"errors"
	"goauth/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância de UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// FindByUsername busca um usuário pelo nome de usuário
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)

	// Retorna nil quando o registro não é encontrado
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Retorna outros erros
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// CreateUser salva um novo usuário no banco de dados
func (r *userRepository) CreateUser(user *models.User) error {
	// Tenta salvar o usuário diretamente
	if err := r.db.Create(user).Error; err != nil {
		// Verifica se o erro é de violação de chave única
		if isUniqueConstraintError(err) {
			return errors.New("username already exists")
		}
		return err
	}

	return nil
}

// isUniqueConstraintError verifica se o erro é de violação de chave única
func isUniqueConstraintError(err error) bool {
	// Aqui você pode implementar verificações específicas dependendo do banco usado (SQLite, Postgres, etc.)
	return false // Ajuste para o seu banco, se necessário
}
