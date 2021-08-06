package grpcv1

import (
	"context"

	v1 "github.com/ku1bu/hourglass/api/hourglass/v1"
	"github.com/ku1bu/hourglass/internal/model"
	"github.com/ku1bu/hourglass/internal/service"

	"github.com/go-kratos/kratos/v2/log"
)

type customer struct {
	v1.UnimplementedCustomerServer
	coordinator *service.Coordinator
	log         *log.Helper
}

func NewCustomer(coordinator *service.Coordinator, logger log.Logger) v1.CustomerServer {
	return &customer{
		coordinator: coordinator,
		log:         log.NewHelper(logger),
	}
}

func (c *customer) Register(in *v1.CustomerRegisterRequest, stream v1.Customer_RegisterServer) error {
	client := new(model.Client)
	client.WatcBack = func(c *model.Client) {
		_ = stream.Send(&v1.CustomerInfo{
			Id:      client.Id,
			Quota:   client.Quota,
			Weights: client.Weights,
		})
	}

	err := c.coordinator.RegisterClient(stream.Context(), in.ProviderName, client)
	if err != nil {
		return err
	}

	<-stream.Context().Done()

	err = c.coordinator.DisregisterClient(stream.Context(), in.ProviderName, client.Id)
	if err != nil {
		return err
	}

	return nil
}

func (c *customer) Feedback(ctx context.Context, in *v1.FeedbackRequest) (*v1.CustomerInfo, error) {
	client, err := c.coordinator.LowActive(ctx, in.Id, in.ServiceName, in.Consume)
	if err != nil {
		return nil, err
	}

	return &v1.CustomerInfo{
		Id:      client.Id,
		Quota:   client.Quota,
		Weights: client.Weights,
	}, nil
}
