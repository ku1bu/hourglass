package grpcv1

import (
	"context"
	v1 "hourglass/api/hourglass/v1"
	"testing"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func TestRegister(t *testing.T) {
	conn, err := grpc.DialInsecure(context.TODO(), grpc.WithEndpoint("localhost:9000"))
	if err != nil {
		t.Fatal(err)
	}

	cli := v1.NewProviderClient(conn)
	resp, err := cli.Register(context.Background(), &v1.ProviderRegisterRequest{
		Capacity:   1000,
		Name:       "111112",
		MinClients: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}
