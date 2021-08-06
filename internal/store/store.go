package store

import (
	"hourglass/internal/model"
)

type Store interface {
	GetOrNewGroup(string) *model.Group
	GetGroup(string) (*model.Group, error)
	RemoveGroup(string) error
}
