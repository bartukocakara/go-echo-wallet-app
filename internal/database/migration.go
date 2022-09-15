package database

import (
	"roof-stack/internal/entity"

	"gorm.io/gorm"
)

func MigrateEntities(db *gorm.DB) {
	db.AutoMigrate( &entity.Wallet{}, &entity.User{}, 
					&entity.Currency{}, &entity.Transaction{})
}