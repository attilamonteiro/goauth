package repository

import (
	"gorm.io/gorm"
)

// Definindo a estrutura Maintenance
type Maintenance struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"not null"`
	CreatedAt   string
	UpdatedAt   string
}

// Repositório de manutenções
type MaintenanceRepository struct {
	db *gorm.DB
}

// NewMaintenanceRepository cria um novo repositório de manutenção
func NewMaintenanceRepository(db *gorm.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}

// Exemplo de método para buscar todas as manutenções
func (r *MaintenanceRepository) GetAll() ([]Maintenance, error) {
	var maintenances []Maintenance
	err := r.db.Find(&maintenances).Error
	return maintenances, err
}
