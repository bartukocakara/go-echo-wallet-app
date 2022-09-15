package repository

import (
	"roof-stack/internal/entity"

	"gorm.io/gorm"
)

type ICommissionRepository interface {
	List() []*entity.Commission
}

type commissionRepository struct {
	connection *gorm.DB
}

func NewCommissionRepository(dbConn *gorm.DB) *commissionRepository{
	return &commissionRepository{ connection: dbConn}
}

func (db *commissionRepository) List() []*entity.Commission {
	var commissions []*entity.Commission  
    db.connection.Find(&commissions) 

    return commissions 
}