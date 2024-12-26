package controllers

import (
	"goauth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MaintenanceController struct {
	db *gorm.DB
}

// NewMaintenanceController cria um novo controller de manutenção
func NewMaintenanceController(db *gorm.DB) *MaintenanceController {
	return &MaintenanceController{db: db}
}

// GetAll retorna todas as manutenções
func (c *MaintenanceController) GetAll(ctx *gin.Context) {
	var maintenances []models.Maintenance

	// Busca todas as manutenções
	if err := c.db.Find(&maintenances).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching maintenances"})
		return
	}

	// Retorna as manutenções em formato JSON
	ctx.JSON(http.StatusOK, maintenances)
}

// Create cria uma nova manutenção
func (c *MaintenanceController) Create(ctx *gin.Context) {
	var maintenance models.Maintenance

	// Faz a bind do JSON para a estrutura de manutenção
	if err := ctx.ShouldBindJSON(&maintenance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Salva a nova manutenção no banco
	if err := c.db.Create(&maintenance).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating maintenance"})
		return
	}

	// Retorna a manutenção criada
	ctx.JSON(http.StatusCreated, maintenance)
}

// Update atualiza uma manutenção existente
func (c *MaintenanceController) Update(ctx *gin.Context) {
	var maintenance models.Maintenance
	id := ctx.Param("id")

	// Busca a manutenção pelo ID
	if err := c.db.First(&maintenance, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Maintenance not found"})
		return
	}

	// Faz a bind dos dados atualizados
	if err := ctx.ShouldBindJSON(&maintenance); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Atualiza os dados da manutenção
	if err := c.db.Save(&maintenance).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating maintenance"})
		return
	}

	// Retorna a manutenção atualizada
	ctx.JSON(http.StatusOK, maintenance)
}

// Delete deleta uma manutenção existente
func (c *MaintenanceController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var maintenance models.Maintenance

	// Busca a manutenção pelo ID
	if err := c.db.First(&maintenance, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Maintenance not found"})
		return
	}

	// Deleta a manutenção
	if err := c.db.Delete(&maintenance).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting maintenance"})
		return
	}

	// Retorna uma mensagem de sucesso
	ctx.JSON(http.StatusOK, gin.H{"message": "Maintenance deleted"})
}
