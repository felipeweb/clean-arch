package repository

import (
	"context"
	"sync"

	"github.com/felipeweb/clean-arch/entity"
)

type InMemory struct {
	m *sync.Map
}

func NewInMemory() *InMemory {
	return &InMemory{
		m: &sync.Map{},
	}
}

func (m *InMemory) Save(_ context.Context, port *entity.Port) error {
	m.m.Store(port.Key, port)
	return nil
}
