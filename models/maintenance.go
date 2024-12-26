package models

type Maintenance struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string
	Description string
}
