package store

import (
	"github.com/ku1bu/hourglass/internal/model"
)

type Store interface {
	GetOrNewGroup(string) *model.Group
	GetGroup(string) (*model.Group, error)
	RemoveGroup(string) error
}
