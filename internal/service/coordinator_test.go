package service

import (
	"context"
	"fmt"
	"hourglass/internal/model"
	"testing"
	"time"
)

type testCoordinator struct {
	coo         *Coordinator
	serviceName string
}

var testCoordinatorData = testCoordinator{
	coo:         NewCoordinator(),
	serviceName: "serviceName",
}

func Test_coordinator_DisregisterService(t *testing.T) {
	err := _coordinator_DisregisterService()
	if err != nil {
		t.Fatal(fmt.Sprintf("%+v", err))
	}

	time.Sleep(time.Second)
}

func _coordinator_DisregisterService() error {
	svc := &model.Service{
		Capacity:   10000,
		Name:       testCoordinatorData.serviceName,
		MinClients: 3,
	}
	err := testCoordinatorData.coo.RegisterService(context.Background(), svc)
	if err != nil {
		return err
	}

	err = testCoordinatorData.coo.RegisterClient(context.Background(), testCoordinatorData.serviceName, &model.Client{
		WatcBack: func(c *model.Client) { fmt.Printf("novel quota:%d, %s\n", c.Quota, c.Id) },
	})
	if err != nil {
		return err
	}

	err = testCoordinatorData.coo.RegisterClient(context.Background(), testCoordinatorData.serviceName, &model.Client{
		WatcBack: func(c *model.Client) { fmt.Printf("novel quota:%d, %s\n", c.Quota, c.Id) },
	})
	if err != nil {
		return err
	}

	err = testCoordinatorData.coo.RegisterClient(context.Background(), testCoordinatorData.serviceName, &model.Client{
		WatcBack: func(c *model.Client) { fmt.Printf("novel quota:%d, %s\n", c.Quota, c.Id) },
	})
	if err != nil {
		return err
	}

	err = testCoordinatorData.coo.RegisterClient(context.Background(), testCoordinatorData.serviceName, &model.Client{
		WatcBack: func(c *model.Client) { fmt.Printf("novel quota:%d, %s\n", c.Quota, c.Id) },
	})
	if err != nil {
		return err
	}

	err = testCoordinatorData.coo.RegisterClient(context.Background(), testCoordinatorData.serviceName, &model.Client{
		WatcBack: func(c *model.Client) { fmt.Printf("novel quota:%d, %s\n", c.Quota, c.Id) },
	})
	if err != nil {
		return err
	}

	c := &model.Client{
		WatcBack: func(c *model.Client) { fmt.Printf("novel quota:%d, %s\n", c.Quota, c.Id) },
	}
	err = testCoordinatorData.coo.RegisterClient(context.Background(), testCoordinatorData.serviceName, c)
	if err != nil {
		return err
	}

	_, err = testCoordinatorData.coo.LowActive(context.Background(), c.Id, testCoordinatorData.serviceName, 0.01)
	if err != nil {
		return err
	}

	err = testCoordinatorData.coo.DisregisterService(context.Background(), svc.Id, svc.Name)
	if err != nil {
		return err
	}

	return nil
}
