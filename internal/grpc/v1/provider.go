package grpcv1

import (
	"context"
	v1 "hourglass/api/hourglass/v1"
	"hourglass/internal/model"
	"hourglass/internal/service"

	"github.com/go-kratos/kratos/v2/log"
)

type provider struct {
	v1.UnimplementedProviderServer
	coordinator *service.Coordinator
	log         *log.Helper
}

func NewProvider(coordinator *service.Coordinator, logger log.Logger) v1.ProviderServer {
	return &provider{coordinator: coordinator, log: log.NewHelper(logger)}
}

func (p *provider) Register(ctx context.Context, in *v1.ProviderRegisterRequest) (*v1.ProviderInfo, error) {
	svc := &model.Service{
		Capacity:   in.Capacity,
		Name:       in.Name,
		MinClients: in.MinClients,
	}

	err := p.coordinator.RegisterService(ctx, svc)
	if err != nil {
		return nil, err
	}

	return &v1.ProviderInfo{
		Id:         svc.Id,
		Capacity:   svc.Capacity,
		Name:       svc.Name,
		MinClients: svc.MinClients,
	}, nil
}

func (p *provider) Disregister(ctx context.Context, in *v1.ProviderDisregisterRequest) (*v1.ProviderInfo, error) {
	err := p.coordinator.DisregisterService(ctx, in.Id, in.Name)
	if err != nil {
		return nil, err
	}

	return &v1.ProviderInfo{Id: in.Id, Name: in.Name}, nil
}
