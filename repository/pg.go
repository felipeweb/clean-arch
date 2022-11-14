package repository

import (
	"context"
	"encoding/json"

	"github.com/felipeweb/clean-arch/entity"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type PortsTable struct {
	Key      string `gorm:"primaryKey"`
	PortInfo datatypes.JSON
}

type PG struct {
	db *gorm.DB
}

func NewPG(db *gorm.DB) *PG {
	db.AutoMigrate(&PortsTable{})
	return &PG{
		db: db,
	}
}

func (p *PG) Save(ctx context.Context, port *entity.Port) error {
	b, err := json.Marshal(port.PortInfo)
	if err != nil {
		return err
	}
	table := &PortsTable{
		Key:      port.Key,
		PortInfo: datatypes.JSON(b),
	}
	return p.db.WithContext(ctx).Save(table).Error
}
