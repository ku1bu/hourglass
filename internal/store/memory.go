package store

import (
	"github.com/ku1bu/hourglass/internal/model"
	"github.com/ku1bu/hourglass/internal/pkg/errorx"
	"sync"
)

var _ Store = new(memoryStore)

type memoryStore struct {
	cache sync.Map
}

func NewMemoryStore() Store {
	return &memoryStore{}
}

func (m *memoryStore) GetOrNewGroup(serviceName string) *model.Group {
	v, ok := m.cache.Load(serviceName)
	if ok {
		return v.(*model.Group)
	}

	v, _ = m.cache.LoadOrStore(serviceName, &model.Group{ClientInfo: make(map[string]*model.Client), Services: make(map[string]*model.Service)})

	return v.(*model.Group)
}

func (m *memoryStore) GetGroup(serviceName string) (*model.Group, error) {
	v, ok := m.cache.Load(serviceName)
	if !ok {
		return nil, errorx.ErrNotFound.WithStack()
	}

	return v.(*model.Group), nil
}

func (m *memoryStore) RemoveGroup(serviceName string) error {
	m.cache.Delete(serviceName)
	return nil
}
