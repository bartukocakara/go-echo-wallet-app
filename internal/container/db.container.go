package container

import (
	config "roof-stack/internal/database"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB = config.SetupDatabaseConnection()
)