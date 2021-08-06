package service

import (
	"context"
	"hourglass/internal/model"
	"hourglass/internal/store"

	"github.com/google/uuid"
)

// 初始通过权重分配 令牌数量
// 当有新节点上下线或LowActive触发重平衡，然后重新分配各客户端令牌数
// 通过长连接watch重平衡后变化

type Coordinator struct {
	store store.Store
}

func NewCoordinator() *Coordinator {
	return &Coordinator{store: store.NewMemoryStore()}
}

func (coo *Coordinator) RegisterService(ctx context.Context, svc *model.Service) error {
	svc.Id = uuid.NewString()

	group := coo.store.GetOrNewGroup(svc.Name)

	group.AddService(svc)
	return nil
}

func (coo *Coordinator) DisregisterService(ctx context.Context, id, serviceName string) error {
	group, err := coo.store.GetGroup(serviceName)
	if err != nil {
		return err
	}

	group.RemoveService(id)
	return nil
}

func (coo *Coordinator) RegisterClient(ctx context.Context, serviceName string, client *model.Client) error {
	client.Id = uuid.NewString()

	group, err := coo.store.GetGroup(serviceName)
	if err != nil {
		return err
	}

	group.AddClient(client)
	return nil
}

func (coo *Coordinator) DisregisterClient(ctx context.Context, groupName, id string) error {
	group, err := coo.store.GetGroup(groupName)
	if err != nil {
		return err
	}

	group.RemoveClient(id)
	return nil
}

func (coo *Coordinator) LowActive(ctx context.Context, id, serviceName string, percentage float32) (*model.Client, error) {
	group, err := coo.store.GetGroup(serviceName)
	if err != nil {
		return nil, err
	}

	group.LowActive(id, percentage)

	return nil, nil
}
