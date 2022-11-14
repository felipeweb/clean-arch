//go:generate mockgen -source=./interface.go -destination=./mock/mock.go -package=mock
package usecase

import (
	"context"

	"github.com/felipeweb/clean-arch/entity"
)

// PortRepository is the repository of ports.
type PortRepository interface {
	Save(ctx context.Context, port *entity.Port) error
}

// PortUsecase is the usecase of ports.
type PortUsecase interface {
	Save(ctx context.Context, key string, info *entity.PortInfo) (*entity.Port, error)
}
