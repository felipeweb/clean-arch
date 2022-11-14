package usecase

import (
	"context"

	"github.com/felipeweb/clean-arch/entity"
)

type PortService struct {
	repo PortRepository
}

func NewPortService(repo PortRepository) *PortService {
	return &PortService{repo: repo}
}

func (s *PortService) Save(ctx context.Context, key string, info *entity.PortInfo) (*entity.Port, error) {
	port := entity.Port{
		Key:      key,
		PortInfo: *info,
	}
	if err := s.repo.Save(ctx, &port); err != nil {
		return nil, err
	}
	return &port, nil
}
