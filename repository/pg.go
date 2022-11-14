package repository

import (
	"context"

	"github.com/felipeweb/clean-arch/entity"
	"gorm.io/gorm"
)

type PG struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *PG {
	db.AutoMigrate(&entity.Port{})
	return &PG{
		db: db,
	}
}

func (p *PG) Save(ctx context.Context, port *entity.Port) error {
	return p.db.WithContext(ctx).Save(port).Error
}
