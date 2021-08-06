package grpcv1

import (
	"context"
	v1 "github.com/ku1bu/hourglass/api/hourglass/v1"
	"testing"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func TestCustoerRegister(t *testing.T) {
	conn, err := grpc.DialInsecure(context.TODO(), grpc.WithEndpoint("localhost:9000"))
	if err != nil {
		t.Fatal(err)
	}

	cli := v1.NewCustomerClient(conn)
	rc, err := cli.Register(context.Background(), &v1.CustomerRegisterRequest{
		ProviderName: "111112",
	})
	if err != nil {
		t.Fatal(err)
	}

	for {
		result, err := rc.Recv()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(result)
	}

}
